package main

import (
	"bytes"
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rjxby/go-htmx-skeleton/backend/app/templates"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/", "../frontend")

	e.GET("/get-example", exampleAppWithHtmx)

	e.Logger.Fatal(e.Start(":8080"))
}

func exampleAppWithHtmx(c echo.Context) error {
	MustExecute := func(tmpl *template.Template, wr io.Writer, data interface{}) {
		if err := tmpl.Execute(wr, data); err != nil {
			panic(err)
		}
	}
	MustRead := func(path string) string {
		file, err := templates.Read(path)
		if err != nil {
			panic(err)
		}
		return string(file)
	}
	tmplstr := MustRead("example.html.tmpl")
	tmpl := template.Must(template.New("example").Parse(tmplstr))
	msg := bytes.Buffer{}
	MustExecute(tmpl, &msg, nil)

	return c.String(http.StatusOK, msg.String())
}
