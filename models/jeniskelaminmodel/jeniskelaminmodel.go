package jeniskelaminmodel

import (
	"log"
	"mahasiswa/config"
	"mahasiswa/entities"
)

func GetAll() []entities.Jeniskelamin {
    rows, err := config.DB.Query(`SELECT * FROM tb_jenis_kelamin`)

    if err != nil {
        log.Println("Error executing query:", err)
        return nil
    }

    defer rows.Close()

    var jeniskelamins []entities.Jeniskelamin

    for rows.Next() {
        var jeniskelaminrow entities.Jeniskelamin

        if err := rows.Scan(&jeniskelaminrow.Id_Jenis_Kelamin, &jeniskelaminrow.Nama_Jenis_kelamin, &jeniskelaminrow.CreatedAt, &jeniskelaminrow.UpdatedAt); err != nil {
            log.Println("Error scanning row:", err)
            return nil
        }

        jeniskelamins = append(jeniskelamins, jeniskelaminrow)
    }

    return jeniskelamins
}
