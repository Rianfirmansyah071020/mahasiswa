package jeniskelamincontroller

import (
	"html/template"
	"mahasiswa/entities"
	"mahasiswa/models/jeniskelaminmodel"
	"net/http"
	"strconv"
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
	
	jeniskelamin.Nama_Jenis_Kelamin = r.FormValue("nama_jenis_kelamin")
	jeniskelamin.CreatedAt = time.Now()
	jeniskelamin.UpdatedAt = time.Now()

	sukses := jeniskelaminmodel.Store(jeniskelamin)	


	if !sukses {
		// create message		
		http.Redirect(w, r, "/jeniskelamin/create", http.StatusFound)
	} else {
		http.Redirect(w, r, "/jeniskelamin/create", http.StatusFound)
	}
	

}


func Delete(w http.ResponseWriter, r *http.Request) { 

	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err.Error())
	}

	if err:= jeniskelaminmodel.Delete(id); err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/jeniskelamin", http.StatusFound)
}