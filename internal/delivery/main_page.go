package delivery

import (
	"net/http"
	"text/template"
)

func HMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ServErrors(w, TextProcessing{http.StatusText(http.StatusNotFound), http.StatusNotFound})
		return
	}
	if r.Method != http.MethodGet {
		ServErrors(w, TextProcessing{http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed})
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ServErrors(w, TextProcessing{http.StatusText(http.StatusNotFound), http.StatusNotFound})
		return
	}
	err1 := tmpl.Execute(w, nil)
	if err1 != nil {
		return
	}
}
