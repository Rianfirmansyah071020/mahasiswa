package web

import (
	"mahasiswa/controllers/homecontroller"
	"net/http"
)

func Route() {
	http.HandleFunc("/", homecontroller.Index)
}