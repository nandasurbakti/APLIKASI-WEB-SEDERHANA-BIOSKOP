package main

import (
	"bioskop_app/config"
	"bioskop_app/database"
	"bioskop_app/routers"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	cfg := config.GetDbConfig()
	
	var (
		DB *sql.DB
		err error
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", 
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Gagal membuka koneksi database: ", err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}
	fmt.Println("Berhasil terhubung ke database")
	database.DBMigrate(DB)

	r := routers.StartServer(DB, cfg)
	fmt.Println("Server berjalan di port 8080")
	r.Run(":8080")
}