package model

import (
	"crypto/md5"
	"encoding/base64"
	"strings"
)

//Link is the representation of http urls
type Link struct {
	Address   string
	Hash      string
	UsedTimes int
}

//NewLink generates a link struct with url's MD5 hash
func NewLink(url string) Link {
	l := new(Link)
	l.Address = url
	md5 := md5.Sum([]byte(url))
	l.Hash = strings.ReplaceAll(strings.ReplaceAll(base64.StdEncoding.EncodeToString(md5[:])[:6], "/", "_"), "+", "-")
	return *l
}

//GetLinkHash returns first 6 characters of the base64 encoded MD5 HASH
func GetLinkHash(url string) string {
	md5 := md5.Sum([]byte(url))
	return strings.ReplaceAll(strings.ReplaceAll(base64.StdEncoding.EncodeToString(md5[:])[:6], "/", "_"), "+", "-")
}
