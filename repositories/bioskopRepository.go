package repositories

import (
	"database/sql"
	"go-gin-postgre/structs"
)

func InsertBioskop(dbParam *sql.DB, bioskop *structs.Bioskop) {
	sqlStatement := `
		INSERT INTO bioskop (Nama, Lokasi, Rating) VALUES ($1, $2, $3)
		RETURNING id, nama, lokasi, rating
	`

	err := dbParam.QueryRow(sqlStatement, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
	if err != nil {
		panic(err)
	}
}

func GetAllBioskop(dbParam *sql.DB) (result []structs.Bioskop, err error) {
	sqlStatement := "SELECT * FROM bioskop"
	rows, err := dbParam.Query(sqlStatement)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var bioskop structs.Bioskop

		err = rows.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
		if err != nil {
			return
		}
		result = append(result, bioskop)
	}
	return
}

func UpdateBioskop(dbParam *sql.DB, bioskop *structs.Bioskop) {
	sqlStatement := "UPDATE bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4"
    err := dbParam.QueryRow(sqlStatement, bioskop.Nama, bioskop.Lokasi, bioskop.Rating, bioskop.ID)
	if err != nil {
		return
	}
}

func DeleteBioskop(dbParam *sql.DB, bioskop *structs.Bioskop) {
	sqlStatement := "DELETE FROM bioskop WHERE id = $1"
    err := dbParam.QueryRow(sqlStatement, bioskop.ID)
	if err != nil {
		return
	}
}