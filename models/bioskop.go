package models

import (
	"database/sql"
	"fmt"
)

type Bioskop struct {
	ID     int64   `json:"id"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}

// membuat tabel database
func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS bioskop (
		id SERIAL PRIMARY KEY,
		nama VARCHAR(100) NOT NULL,
		lokasi VARCHAR(100) NOT NULL,
		rating FLOAT
	);
	`
    _, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("gagal membuat tabel bioskop: %w", err)
	}
	return nil
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
	rows, err := db.Query(`SELECT id, nama, lokasi, rating FROM bioskop`)
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