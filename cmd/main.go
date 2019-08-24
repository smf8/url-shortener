package main

import (
	"github.com/smf8/url-shortener/db"
	"github.com/smf8/url-shortener/model"
)

func main() {
	db.CreateDB("urls")
	db.AddLink(model.NewLink("http://google.com"))
	defer db.Close()
}
