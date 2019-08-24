package handler

import (
	"github.com/smf8/url-shortener/db"
	"github.com/smf8/url-shortener/model"
	"net/http"
	"strings"
)

//RegisterURL reads a url address and stores it returning it's shorten form
func RegisterURLHandler(wr http.ResponseWriter, r *http.Request) {
	// read url from POST
	url := r.PostFormValue("url")
	if url == "" {
		http.Error(wr, "Url musn't be empty", http.StatusBadRequest)
		return
	}
	link := model.NewLink(url)
	go func() {
		db.AddLink(link)
	}()
}

func RedirectHandler(wr http.ResponseWriter, r *http.Request) {
	// extract url hash from url path
	hash := strings.Trim(strings.TrimSpace(r.URL.Path), "/")
	//fmt.Println(hash)
	l := db.GetLink(hash)
	if l.Address != "" {
		go func() {
			db.IncrementUsage(hash)
		}()
		http.Redirect(wr, r, l.Address, http.StatusSeeOther)
	} else {
		http.Error(wr, "Invalid Url", http.StatusBadRequest)
	}
}
