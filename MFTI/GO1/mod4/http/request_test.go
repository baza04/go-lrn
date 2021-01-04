package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	ID         string
	Response   string
	StatusCode int
}

// GetUser just return status request execution and incoming data
func GetUser(w http.ResponseWriter, r *http.Request) {
	// get value of "id" parameter from request
	key := r.FormValue("id")
	if key == "42" {
		w.WriteHeader(http.StatusOK) // set status 200
		io.WriteString(w, `{"status": 200, "resp": {"user": 42}}`)
	} else {
		w.WriteHeader(http.StatusInternalServerError) // set status 500
		io.WriteString(w, `{"status": 500, "err": "db_error"}`)
	}
}

// TestGetUser do test
func TestGetUser(t *testing.T) {
	// init test cases
	cases := []TestCase{
		TestCase{
			ID:         "42",
			Response:   `{"status": 200, "resp": {"user": 42}}`,
			StatusCode: http.StatusOK,
		},
		TestCase{
			ID:         "500",
			Response:   `{"status": 500, "err": "db_error"}`,
			StatusCode: http.StatusInternalServerError,
		},
	}

	// range arr with test cases
	for caseNum, item := range cases {
		// add ID from case
		url := "http://example.com/api/user?id=" + item.ID
		// init custom req with url above
		req := httptest.NewRequest("GET", url, nil)

		// init testing response object
		w := httptest.NewRecorder()

		// invoking the test handler
		GetUser(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}

		// get value of responce and read
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)
		if bodyStr != item.Response {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyStr, item.Response)
		}
	}
}
