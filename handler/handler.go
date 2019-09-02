package handler

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/smf8/url-shortener/db"
	"github.com/smf8/url-shortener/model"
	"net/http"
	"strings"
)

//SaveUrl gets url form value and saves it inside database
func SaveUrl(c echo.Context) error {
	var err error
	url := c.FormValue("url")
	if url == "" {
		err = errors.New("Url musn't be empty")
		return err
	}
	// remove www. from the beginning if exists
	url = strings.Replace(url, "www.", "", 1)
	if !strings.Contains(url, "http://") {
		url = "http://" + url
	}
	// create a link struct from url
	link := model.NewLink(url)
	go func() {
		db.AddLink(link)
	}()
	err = c.String(http.StatusOK, "Your shortened link is : "+c.Request().Host+"/"+link.Hash)
	return err
}

//MainPage is for showing default form for getting links from user
func MainPage(c echo.Context) error {
	err := c.Render(http.StatusOK, "addlink.html", nil)
	return err
}

//Redirect redirects requests sent by client
func Redirect(c echo.Context) error {
	var err error
	if c.Param("hash") != "" {
		hash := c.Param("hash")
		l := db.GetLink(hash)
		if l.Address != "" {
			// Increment usage field of link in `links` table by 1
			go func() {
				db.IncrementUsage(hash)
			}()
			err = c.Redirect(http.StatusTemporaryRedirect, l.Address)
		} else {
			err = c.String(http.StatusBadRequest, "Invalid url")
		}
	}
	return err
}
