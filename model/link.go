package model

import (
	"crypto/md5"
	"encoding/base64"
)

//Link is the representation of http urls
type Link struct {
	url  string
	hash string
}

//NewLink generates a link struct with url's MD5 hash
func NewLink(url string) {
	l := new(Link)
	l.url = url
	md5 := md5.Sum([]byte(url))
	hash := base64.StdEncoding.EncodeToString(md5[:])
	l.hash = hash[:6]
}
