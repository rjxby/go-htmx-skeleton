package main

import (
	"bytes"
	"embed"
	"io"
	"io/fs"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rjxby/go-htmx-skeleton/backend/app/templates"
)

//go:embed app
var embededFiles embed.FS

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(embededFiles, "app/public")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	assetHandler := http.FileServer(getFileSystem())
	e.GET("/*", echo.WrapHandler(assetHandler))

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
