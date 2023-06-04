package internal

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	if status == http.StatusNotFound {
		html, err := template.ParseFiles("./templates/404.html")
		err = html.Execute(w, nil)
		if err != nil {
			fmt.Println("Cant execute 404 page")
			return
		}
	}
}
