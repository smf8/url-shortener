package main

import (
	"github.com/smf8/url-shortener/db"
	"github.com/smf8/url-shortener/handler"
	"log"
	"net/http"
)

func main() {
	db.CreateDB("urls")
	mux := http.NewServeMux()
	mux.HandleFunc("/new", handler.RegisterURLHandler)
	mux.HandleFunc("/", handler.MainHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
func loadTmpl(filename string) {

}
