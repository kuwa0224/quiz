package main

import (
	"html/template"
	"io"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/labstack/echo"
)

type (
	Template struct {
		templates *template.Template
	}
)

func init() {

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*")),
	}
	e.Renderer = t

	e.GET("/", top)
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		gaec := appengine.NewContext(c.Request())
		log.Errorf(gaec, "failed to execute template: %v", err)
	}

	return err
}

func top(c echo.Context) error {
	gaec := appengine.NewContext(c.Request())

	log.Infof(gaec, "OK")

	return c.Render(http.StatusOK, "top", "Takeshi")
}
