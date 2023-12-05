package jeniskelamincontroller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("views/pages/jeniskelamin/index.html")

	if err != nil {
		panic(err.Error())
	}
	
	
	temp.Execute(w, nil)
}

func Create(w http.ResponseWriter, r *http.Request) { 

	temp, err := template.ParseFiles("views/pages/jeniskelamin/create.html")

	if err != nil {
		panic(err.Error())
	}


	temp.Execute(w, nil)
}