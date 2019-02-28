package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/raion314/goweb/utils"
)

var config utils.Configuration

func init() {
	loadConfig()
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	utils.Logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	utils.P("ChitChat", utils.Version(), "started at", config.Address)

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", utils.Index)
	mux.HandleFunc("/err", utils.Err)

	mux.HandleFunc("/login", utils.Login)
	mux.HandleFunc("/logout", utils.Logout)
	mux.HandleFunc("/signup", utils.Signup)
	mux.HandleFunc("/signup_account", utils.SignupAccount)
	mux.HandleFunc("/authenticate", utils.Authenticate)

	mux.HandleFunc("/thread/new", utils.NewThread)
	mux.HandleFunc("/thread/create", utils.CreateThread)
	mux.HandleFunc("/thread/post", utils.PostThread)
	mux.HandleFunc("/thread/read", utils.ReadThread)

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = utils.Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
