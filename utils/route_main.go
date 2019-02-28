package utils

import (
	"net/http"

	"github.com/raion314/goweb/data"
)

// Err comment
func Err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

// Index comment
func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		errorMessage(w, r, "Cannot get threads")
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		}
	}
}
