package models

import (
	"database/sql"
)

type Bioskop struct {
	ID     int64   `json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}

func InsertData(db *sql.DB, b *Bioskop) error {
	query := `INSERT INTO bioskop (nama, lokasi, rating) 
			  VALUES ($1, $2, $3) RETURNING id`
	
	err := db.QueryRow(query, b.Nama, b.Lokasi, b.Rating).Scan(&b.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBioskop(db *sql.DB, b *Bioskop) error {
	query := `UPDATE bioskop 
			  SET nama=$1, lokasi=$2, rating=$3 
			  WHERE id=$4`
	_, err := db.Exec(query, b.Nama, b.Lokasi, b.Rating, b.ID)
	return err
}

func DeleteBioskop(db *sql.DB, id int64) error {
	query := `DELETE FROM bioskop WHERE id=$1`
	_, err := db.Exec(query, id)
	return err
}

func GetAllBioskop(db *sql.DB) ([]Bioskop, error) {
	query := `SELECT * FROM bioskop`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bioskops []Bioskop
	for rows.Next() {
		var b Bioskop
		if err := rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating); err != nil {
			return nil, err
		}
		bioskops = append(bioskops, b)
	}

	return bioskops, nil
}

func GetBioskopByID(db *sql.DB, id int64) (Bioskop, error) {
	var b Bioskop
	query := "SELECT id, nama, lokasi, rating FROM bioskop WHERE id=$1"
	row := db.QueryRow(query, id)
	err := row.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
	return b, err
}
