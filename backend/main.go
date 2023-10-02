package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/go-chi/render"

	"github.com/rjxby/go-htmx-skeleton/backend/app/templates"
)

func main() {
	http.HandleFunc("/get-example", exampleAppWithHtmx)
	http.Handle("/", http.FileServer(http.Dir("../frontend")))
	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", nil)
}

func exampleAppWithHtmx(w http.ResponseWriter, r *http.Request) {
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
	render.HTML(w, r, msg.String())
}
