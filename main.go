package main

import (
	"log"
	"mahasiswa/routes/web"
	"net/http"
)

func main() {
	// menangani file statis
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// route web
	web.Route()

	// menjalankan server
	log.Println("Server running on port 9999")
	log.Fatal(http.ListenAndServe(":9999", nil))

}