package filestatis

import (
	"net/http"
)

// menangani file statis
func Route() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}