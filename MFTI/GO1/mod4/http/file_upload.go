package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var uploadFormTmpl = []byte(`
<html>
	<body>
	<form action="/upload" method="post" enctype="multipart/form-data">
		Image: <input type="file" name="my_file">
		<input type="submit" value="Upload">
	</form>
	</body>
</html>
`)

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write(uploadFormTmpl)
}

func uploadPage(w http.ResponseWriter, r *http.Request) {
	// parse file from form by memory limit other part of file go to temp dir in pc
	// golang not parse files from form by default, we need to do it by our self
	r.ParseMultipartForm(5 * 1024 * 1025)

	// FormFile data and headers(like name, size)
	file, handler, err := r.FormFile("my_file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // very important

	// print on page file_name from headers of file
	fmt.Fprintf(w, "handler.Filename %v\n", handler.Filename)
	// print mime_type or else headers
	fmt.Fprintf(w, "handler.Header %#v\n", handler.Header)

	// create new object from crypto.md5 pkg
	hasher := md5.New()
	// do hash of readed data
	io.Copy(hasher, file)

	// we can count hash by method of md5 object (now we copy value to object, than call)
	// also hasher.Sum() can generate hash if we put value to method
	fmt.Fprintf(w, "md5 %x\n", hasher.Sum(nil))
}

// Params struct to encoding income JSON
type Params struct {
	ID   int
	User string
}

// to check work of uploadRawBody func use next command:
/*
curl -v -X POST -H "Content-Type: application/json" -d '{"id": 2, "user": "rvasily"}' http://localhost:8080/raw_body
*/

func uploadRawBody(w http.ResponseWriter, r *http.Request) {
	// read full body of incoming request
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// create obj, parse json
	p := &Params{}
	err = json.Unmarshal(body, p)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "content-type %#v\n", r.Header.Get("Content-Type"))
	fmt.Fprintf(w, "params %#v\n", p)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootPage)
	mux.HandleFunc("/upload", uploadPage)
	mux.HandleFunc("/raw_body", uploadRawBody)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("start server at: 8080")
	server.ListenAndServe()
}
