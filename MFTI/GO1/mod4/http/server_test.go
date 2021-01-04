package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type TestCase struct {
	ID      string
	Result  *CheckoutResult
	IsError bool
}

type CheckoutResult struct {
	Status  int
	Balance int
	Err     string
}

// CheckoutDummy send response and set status to incoming request
func CheckoutDummy(w http.ResponseWriter, r *http.Request) {
	// get value from income request
	key := r.FormValue("id")
	// check case
	switch key {
	case "42":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 200, "balance": 100500}`)
	case "100500":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 400, "err": "bad_balance"}`)
	case "__broken_json":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 400`) //broken json
	case "__internal_error":
		fallthrough
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type Cart struct {
	PaymentApiURL string // payment getaway
}

func (c *Cart) Checkout(id string) (*CheckoutResult, error) {
	// init url with recieved id
	url := c.PaymentApiURL + "?id=" + id
	// get response after GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// init object to unmarshaling
	result := &CheckoutResult{}

	// do unmarshaling
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func TestCartCheckout(t *testing.T) {
	cases := []TestCase{
		TestCase{
			ID: "42",
			Result: &CheckoutResult{
				Status:  200,
				Balance: 100500,
				Err:     "",
			},
			IsError: false,
		},
		TestCase{
			ID: "100500",
			Result: &CheckoutResult{
				Status:  400,
				Balance: 0,
				Err:     "bad_balance",
			},
			IsError: false,
		},
		TestCase{
			ID:      "__broken_json",
			Result:  nil,
			IsError: true,
		},
		TestCase{
			ID:      "__internal_error",
			Result:  nil,
			IsError: true,
		},
	}

	// init test server with random port and url CheckoutDummy as HandlerFunc
	ts := httptest.NewServer(http.HandlerFunc(CheckoutDummy))

	// range test cases
	for caseNum, item := range cases {
		// put test server url to new object of Cart
		c := &Cart{
			PaymentApiURL: ts.URL,
		}

		// get result of handling
		result, err := c.Checkout(item.ID)

		if err != nil && !item.IsError {
			t.Errorf("[%d] unexpected error: %#v", caseNum, err)
		}
		if err == nil && item.IsError {
			t.Errorf("[%d] expected error, got nil", caseNum)
		}
		if !reflect.DeepEqual(item.Result, result) {
			t.Errorf("[%d] wrong result, expected %#v, got %#v", caseNum, item.Result, result)
		}
	}
	// close test server
	ts.Close()
}
