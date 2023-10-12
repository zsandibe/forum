package pkg

import (
	"html/template"
	"net/http"
	"log"
	modelsData "forum/internal/models/data"
)



func Error(w http.ResponseWriter, status int, err error) {
	var errMsg string = http.StatusText(status)
	if err != nil {
		if status == http.StatusInternalServerError {
			log.Println(err)
		} else if status != http.StatusNotFound {
			errMsg = err.Error()
		}
	}

	w.WriteHeader(status)

	data := modelsData.Data{
		Error: modelsData.ErrorMsg{
			Status: status,
			Message:    errMsg,
		},
	}

	tmpl,err := template.ParseFiles("./ui/templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	tmpl.Execute(w, data)
}
