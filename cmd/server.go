package main

import (
	"github.com/labstack/echo"
	"github.com/smf8/url-shortener/db"
	"github.com/smf8/url-shortener/handler"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	err := db.NewDB("links")
	if err != nil {
		panic(err)
	}
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("cmd/*.html")),
	}
	e.Renderer = t
	e.GET("/", handler.MainPage)
	e.POST("/new", handler.SaveUrl)
	e.GET("/:hash", handler.Redirect)
	e.Logger.Fatal(e.Start(":8080"))
}
