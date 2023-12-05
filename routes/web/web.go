package web

import (
	"mahasiswa/controllers/homecontroller"
	"mahasiswa/controllers/jeniskelamincontroller"
	"net/http"
)

func Route() {
	http.HandleFunc("/", homecontroller.Index)
	http.HandleFunc("/jeniskelamin", jeniskelamincontroller.Index)
}