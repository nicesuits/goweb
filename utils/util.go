package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/raion314/goweb/data"
)

// Configuration struct
type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var Logger *log.Logger

// P comment
func P(a ...interface{}) {
	fmt.Println(a...)
}

func errorMessage(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{UUID: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t := template.Must(template.ParseFiles(files...))
	t.ExecuteTemplate(w, "layout", data)
}

// Config comment
func Config() (config Configuration) {
	return config
}

// Info comment
func Info(args ...interface{}) {
	Logger.SetPrefix("INFO")
	Logger.Println(args...)
}

// Danger comment
func Danger(args ...interface{}) {
	Logger.SetPrefix("ERROR")
	Logger.Println(args...)
}

// Warning comment
func Warning(args ...interface{}) {
	Logger.SetPrefix("WARNING")
	Logger.Println(args...)
}

// Version comment
func Version() string {
	return "0.1.0"
}
