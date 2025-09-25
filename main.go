package main

import (
	"bioskop_app/config"
	"bioskop_app/controllers"
	"bioskop_app/models"
	"bioskop_app/routers"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Config database
	dbConfig := config.GetDbConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Gagal membuka koneksi database: ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}

	fmt.Println("Berhasil terhubung ke database")
	err = models.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}

	// assign db ke controller global
	controllers.DB = db

	// tabel database
	if err := models.CreateTable(db); err != nil {
		log.Fatal("Gagal bikin tabel:", err)
	}

	r := routers.StartServer()
	r.Run(":8080")
}