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

        if err := rows.Scan(&jeniskelaminrow.Id_Jenis_Kelamin, &jeniskelaminrow.Nama_Jenis_Kelamin, &jeniskelaminrow.CreatedAt, &jeniskelaminrow.UpdatedAt); err != nil {
            log.Println("Error scanning row:", err)
            return nil
        }

        jeniskelamins = append(jeniskelamins, jeniskelaminrow)
    }

    return jeniskelamins
}


func Store(jeniskelamin entities.Jeniskelamin) bool {

    temp, err := config.DB.Exec(`INSERT INTO tb_jenis_kelamin (id_jenis_kelamin, nama_jenis_kelamin, created_at, updated_at) VALUES (?, ?, ?, ?)`, jeniskelamin.Id_Jenis_Kelamin, jeniskelamin.Nama_Jenis_Kelamin, jeniskelamin.CreatedAt, jeniskelamin.UpdatedAt)

    if err != nil {
        log.Println("Error executing query:", err)
        return false
    }

    _, err = temp.LastInsertId()

    if err != nil {
        log.Println("Error getting last insert id:", err)
        return false
    }

    return true
}


func Delete(id int) error {
    _,err := config.DB.Exec(`DELETE FROM tb_jenis_kelamin WHERE id_jenis_kelamin = ?`, id)
    if err != nil {
        log.Println("Error executing query:", err)
        return err
    }

    return nil
}


func Edit(id int) entities.Jeniskelamin {

   rows := config.DB.QueryRow(`SELECT id_jenis_kelamin, nama_jenis_kelamin, created_at, updated_at FROM tb_jenis_kelamin WHERE id_jenis_kelamin = ?`, id)

   var jeniskelamin entities.Jeniskelamin

   if err := rows.Scan(&jeniskelamin.Id_Jenis_Kelamin, &jeniskelamin.Nama_Jenis_Kelamin, &jeniskelamin.CreatedAt, &jeniskelamin.UpdatedAt); err != nil {
       log.Println("Error scanning row:", err)
       return jeniskelamin
   }

   return jeniskelamin


}



func Update(id int, jeniskelamin entities.Jeniskelamin) bool {

    _, err := config.DB.Exec(`UPDATE tb_jenis_kelamin SET nama_jenis_kelamin = ?, updated_at = ? WHERE id_jenis_kelamin = ?`, jeniskelamin.Nama_Jenis_Kelamin, jeniskelamin.UpdatedAt, id)

    if err != nil {
        log.Println("Error executing query:", err)
        return false
    }

    return true
}