package main

import (
	"log"
	"mahasiswa/config"
	"mahasiswa/routes/filestatis"
	"mahasiswa/routes/web"
	"net/http"
)

func main() {

	config.ConnectDB()
	
	// menjalankan file statis handle
	filestatis.Route()

	// route web
	web.Route()


	// menjalankan server
	log.Println("Server running on port 9999")
	log.Fatal(http.ListenAndServe(":9999", nil))

}