package handler

import (
	"net/http"
	"text/template"

	"github.com/tnnz20/Scalable-Web-Service-with-Golang/pkg/helpers"
)

type Data struct {
	Water       int
	Wind        int
	WaterStatus string
	WindStatus  string
}

func RenderHTML(w http.ResponseWriter, r *http.Request) {

	element, err := helpers.LoadJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status := helpers.CheckElement(element)

	data := Data{
		Water:       int(element.Water),
		Wind:        int(element.Wind),
		WaterStatus: status.WaterStatus,
		WindStatus:  status.WindStatus,
	}

	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
