package controllers

import (
	"bioskop_app/models"
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func CreateBioskop(c *gin.Context) {
	var b models.Bioskop
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if strings.TrimSpace(b.Nama) == "" || strings.TrimSpace(b.Lokasi) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
			return
		}

	if err := models.InsertData(DB, &b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal tambah data"})
		return
	}

	c.JSON(http.StatusCreated, b)
}

func UpdateBioskop(c *gin.Context) {
	var b models.Bioskop
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateBioskop(DB, &b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diupdate"})
}

func DeleteBioskop(c *gin.Context) {
	var input struct {
		ID int64 `json:"id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DeleteBioskop(DB, input.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal hapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}

func GetBioskop(c *gin.Context) {
	bioskops, err := models.GetAllBioskop(DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal ambil data"})
		return
	}

	c.JSON(http.StatusOK, bioskops)
}