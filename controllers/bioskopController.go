package controllers

import (
	"go-gin-postgre/database"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Bioskop struct {
	ID		int		`json:"id"`
	Nama	string	`json:"nama"`
	Lokasi	string	`json:"lokasi"`
	Rating	float64	`json:"rating"`
}

func CreateBioskop(ctx *gin.Context) {
	db, _ := database.GetConnection()
	defer db.Close()

	var bioskop Bioskop
	if err := ctx.ShouldBindJSON(&bioskop); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	errorList := map[string]string{}


	if bioskop.Nama == "" {
		errorList["nama"] = "Nama is required"
	}

	if bioskop.Lokasi == "" {
		errorList["lokasi"] = "Lokasi is required"
	}

	if len(errorList) > 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errorList,
		})
		return
	}

	sqlStatement := `
	INSERT INTO bioskop (Nama, Lokasi, Rating) VALUES ($1, $2, $3) Returning *
	`

	err := db.QueryRow(sqlStatement, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"bioskop": bioskop,
	})
}