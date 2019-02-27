package main

import (
	"fmt"
	"net/http"

	"github.com/raion314/goweb/data"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout")
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup")
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup_account")
}
