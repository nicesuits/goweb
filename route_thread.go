package main

import (
	"fmt"
	"net/http"
)

func newThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}

func createThread(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createThread")
}

func readThread(w http.ResponseWriter, r *http.Request) {
	fmt.Println("readThread")

}

func postThread(w http.ResponseWriter, r *http.Request) {
	fmt.Println("postThread")

}
