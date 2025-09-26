package controllers

import (
	"bioskop_app/database"
	"bioskop_app/models"
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)



func PostBioskop(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
	var b models.Bioskop
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if strings.TrimSpace(b.Nama) == "" || strings.TrimSpace(b.Lokasi) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
			return
		}

	if err := models.InsertData(database.DB, &b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal tambah data"})
		return
	}

	c.JSON(http.StatusCreated, b)
	}
	
}

func UpdateBioskop(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
	var b models.Bioskop
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateBioskop(database.DB, &b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diupdate"})
	}
}


func DeleteBioskop(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
	var input struct {
		ID int64 `json:"id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DeleteBioskop(database.DB, input.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal hapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
	}
}

func GetBioskop(c *gin.Context) {
	bioskops, err := models.GetAllBioskop(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal ambil data"})
		return
	}

	c.JSON(http.StatusOK, bioskops)
}

func GetBioskopID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	bioskop, err := models.GetBioskopByID(database.DB, id)
	if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data"})
        }
        return
    }

	c.JSON(http.StatusOK, bioskop)
}