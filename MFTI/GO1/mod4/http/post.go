package main

import (
	"fmt"
	"net/http"
)

var loginFormTmpl = []byte(`
<html>
	<body>
	<form action="/" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	</body>
</html>
`)

func mainPage(w http.ResponseWriter, r *http.Request) {
	// will print login tmpl if method is not POST
	if r.Method != http.MethodPost {
		w.Write(loginFormTmpl)
		return
	}

	// if method is POST we will work with request parameters
	// we can get data from request by ourself
	// or user func FormValue

	/*
		r.ParseForm()  						// Parse form
		inputLogin := r.Form["login"][0]	// get data from parsed form
	*/

	// FormValue return data from all methods GET or POST
	// but POST method have more priority
	inputLogin := r.FormValue("login")
	fmt.Fprintf(w, "you enter: %s\n", inputLogin)
}
func main() {
	http.HandleFunc("/", mainPage)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
