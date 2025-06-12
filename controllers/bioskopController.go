package controllers

import (
	"go-gin-postgre/database"
	"go-gin-postgre/repositories"
	"go-gin-postgre/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db, _ = database.GetConnection()

func CreateBioskop(ctx *gin.Context) {
	var bioskop structs.Bioskop
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

	repositories.InsertBioskop(db, &bioskop)

	ctx.JSON(http.StatusOK, gin.H{
		"bioskop": bioskop,
	})
}

func GetAllBioskop(ctx *gin.Context) {
	bioskop, err := repositories.GetAllBioskop(db)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"bioskop": bioskop,
	})
}

func UpdateBioskop(ctx *gin.Context) {
	var bioskop structs.Bioskop
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&bioskop); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bioskop.ID = id

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

	repositories.UpdateBioskop(db, &bioskop)

	ctx.JSON(http.StatusOK, gin.H{
		"bioskop": bioskop,
	})
}

func DeleteBioskop(ctx *gin.Context) {
	var bioskop structs.Bioskop
	id, _ := strconv.Atoi(ctx.Param("id"))

	bioskop.ID = id
	repositories.DeleteBioskop(db, &bioskop)
	var message string = "Bioskop with id " + strconv.Itoa(id) + " succesfully deleted"
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}