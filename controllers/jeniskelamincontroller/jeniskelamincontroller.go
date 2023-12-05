package jeniskelamincontroller

import (
	"html/template"
	"mahasiswa/entities"
	"mahasiswa/models/jeniskelaminmodel"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {

	jeniskelamins := jeniskelaminmodel.GetAll()

	data := map[string]interface{}{
		"jeniskelamins": jeniskelamins,
	}

	temp, err := template.ParseFiles("views/pages/jeniskelamin/index.html")

	if err != nil {
		panic(err.Error())
	}
	
	
	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) { 

	temp, err := template.ParseFiles("views/pages/jeniskelamin/create.html")

	if err != nil {
		panic(err.Error())
	}


	temp.Execute(w, nil)
}

func Store(w http.ResponseWriter, r *http.Request) { 

	var jeniskelamin entities.Jeniskelamin
	
	jeniskelamin.Nama_Jenis_kelamin = r.FormValue("nama_jenis_kelamin")
	jeniskelamin.CreatedAt = time.Now()
	jeniskelamin.UpdatedAt = time.Now()

}