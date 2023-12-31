package web

import (
	"mahasiswa/controllers/homecontroller"
	"mahasiswa/controllers/jeniskelamincontroller"
	"net/http"
)

func Route() {
	http.HandleFunc("/", homecontroller.Index)
	http.HandleFunc("/jeniskelamin", jeniskelamincontroller.Index)
	http.HandleFunc("/jeniskelamin/create", jeniskelamincontroller.Create)
	http.HandleFunc("/jeniskelamin/store", jeniskelamincontroller.Store)
	http.HandleFunc("/jeniskelamin/delete", jeniskelamincontroller.Delete)
	http.HandleFunc("/jeniskelamin/edit", jeniskelamincontroller.Edit)
	http.HandleFunc("/jeniskelamin/update", jeniskelamincontroller.Update)
}