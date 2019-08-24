package handler

import (
	"github.com/smf8/url-shortener/db"
	"github.com/smf8/url-shortener/model"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func init() {
	db.CreateDB("tests")
	mux := http.NewServeMux()
	mux.HandleFunc("/new", RegisterURLHandler)
	mux.HandleFunc("/", RedirectHandler)
	go func() { log.Fatal(http.ListenAndServe(":3030", mux)) }()
}
func TestRegister(t *testing.T) {
	l := url.Values{}
	l.Set("url", "http://github.com")
	res, err := http.PostForm("http://localhost:3030/new", l)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Error("Invalid response code ", res.StatusCode)
	}
	j := db.GetLink(model.GetLinkHash("http://github.com"))
	if j.Address != "http://github.com" {
		t.Error("Link Not found in database", "s")
	}
	//defer os.Exit(0)
}
func TestRedirect(t *testing.T) {
	l := model.GetLinkHash("http://github.com")
	resp, err := http.Get("http://localhost:3030/" + l)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Error on redirection", resp.Status)
	}
	defer os.Remove("tests.db")
}
