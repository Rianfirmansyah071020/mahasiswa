package homecontroller

import (
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err :=	template.ParseFiles("views/pages/home/index.html")

	if err != nil {
		panic(err.Error())
	}

	temp.Execute(w, nil)
}