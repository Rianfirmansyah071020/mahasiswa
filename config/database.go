package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/mahasiswa?parseTime=true")

	if err != nil {
		log.Fatal("Error connecting to database:", err)
		return
	}

	err = db.Ping()
	if err != nil {
		db.Close() // Tutup koneksi jika ada kesalahan ping
		log.Fatal("Error pinging database:", err)
		return
	}

	DB = db
	
	log.Println("Connected to database")
}
