package utils

import (
	"fmt"
	"net/http"

	"github.com/raion314/goweb/data"
)

// Authenticate comment
func Authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		Danger(err, "Cannot find user")
	}
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			Danger(err, "Cannot create session")
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

// Login comment
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
}

// Logout comment
func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout")
}

// Signup comment
func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup")
}

// SignupAccount comment
func SignupAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup_account")
}
