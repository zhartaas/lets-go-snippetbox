package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // changed
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/footer.partial.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err) // change there
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.serverError(w, err) // and there
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w) // change
		return
	}

	fmt.Fprintf(w, "Displaying a specific snippet %d\n", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed) // change
		return
	}

	w.Write([]byte("create snippet"))
}
