package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "getHandler: incoming request %#v\n", r)
		fmt.Fprintf(w, "getHandler: r.URL %#v\n", r.URL)
	})

	http.HandleFunc("raw_body", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Fprint(w, "postHandler: raw body %s\n", string(body))
	})

	fmt.Println("starting server at: 8080")
	http.ListenAndServe(":8080", nil)
}

// simple
func runGet() {
	url := "http://127.0.0.1:8080/?param=123&param2=test"
	// use Get func to do simple GET request with choosen URL
	// http.Get return response object
	resp, err := http.Get(url) // we can use response to return some answer
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer resp.Body.Close() // very important

	// just to see entire of response
	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("http.Get body %#v\n\n\n", string(respBody))
}

// full request with custom parameters
func runGetFullReq() {
	// create new custom request object
	req := &http.Request{
		Method: http.MethodGet, // set method
		Header: http.Header{ // set headers
			"User-Agent": {"cousera-golang"},
		},
	}
	// we can also set url for request, but we do it by url.Parse func below

	// parse url and add parsed url to created request object
	req.URL, _ = url.Parse("http://127.0.0.1:8080/?id=42")
	// add "query parameter" to request object in "url" field
	req.URL.Query().Set("user", "Make")

	// do request with created request object parameters
	// and !!! "DEFAULT CLIENT" !!!
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	// close request body
	defer resp.Body.Close() // !!! very important

	respBody, err := ioutil.ReadAll(resp.Body) // read resp body just for printin in log
	fmt.Printf("testGetFullReq resp %#v\n\n\n", string(respBody))
}

func runTransportAndPost() {

	// create transport custom object
	// use to custom client object
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	// init and parse data to request
	data := `{"id": 42, "user": "Make"}`
	body := bytes.NewBufferString(data)

	url := "http://127.0.0.1:8080/raw_body" // init url to req
	// init new request object with POST method, url and parse data
	req, _ := http.NewRequest(http.MethodPost, url, body)
	req.Header.Add("Content-Type", "application/json")        // set Content-Type header
	req.Header.Add("Content-Length", strconv.Itoa(len(data))) // set Content-Length header

	// do request with created request object parameters
	// and !!! "CUSTOM CLIENT" !!!
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer resp.Body.Close()

	// read only for printing in log
	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("runTransport %#v\n\n\n", string(respBody))
}

func main() {
	// start server in goroutine
	go startServer()

	time.Sleep(100 * time.Millisecond)

	runGet()        // do simple GET request
	runGetFullReq() // do custom GET request

	// do custom POST request with custom Client and Transport
	runTransportAndPost()
}
