package main

import (
	"fmt"
	"net/http"
	"time"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	// get Cookie from request
	session, err := r.Cookie("session_id")
	// check logged status
	loggedIn := (err != http.ErrNoCookie)

	if loggedIn { // if logged show link to logout
		fmt.Fprintf(w, `<a href="/logout">logout</a>`)
		fmt.Fprintf(w, "Welcome, %v", session.Value)
	} else { // if not show link to login
		fmt.Fprintf(w, `<a href="/login">login</a>`)
		fmt.Fprintf(w, "You need to login")
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	// init expiration time for cookie when logged
	expiration := time.Now().Add(10 * time.Hour)
	// create Cookie
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    "baza04",
		Expires:  expiration,
		HttpOnly: true,
	}
	// set cookie
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func logoutPage(w http.ResponseWriter, r *http.Request) {
	// get cookie by name
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)

	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/logout", logoutPage)
	http.HandleFunc("/", mainPage)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
