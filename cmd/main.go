package main

import (
	"fmt"
	"github.com/smf8/url-shortener/db"
	"github.com/smf8/url-shortener/model"
)

func main() {
	db.CreateDB("urls")
	db.AddLink(model.NewLink("http://google.com"))
	l := db.GetLink(model.GetLinkHash("http://google.com"))
	fmt.Println(l.Address)
	defer db.Close()
}
