package controllers

import (
	"belajar-gin/models"
	"fmt"
	"net/http"

	"belajar-gin/database"

	"github.com/gin-gonic/gin"
)

// type Car struct {
// 	CarID string `json:"car_id"`
// 	Brand string `json:"brand"`
// 	Model string `json:"model"`
// 	Price int    `json:"price"`
// }

var CarDatas = []models.Car{}

func CreateCar(ctx *gin.Context) {
	db := database.GetDB()

	Car := models.Car{
		Brand: brand,
		Model: model,
		Price: price,
	}

	err := db.Create(&Car).Error

	if err != nil {
		fmt.Println("error creating Car data:", err)
		return
	}

	fmt.Println("new Car data:", Car)

	var newCar models.Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

// func CreateCar(ctx *gin.Context) {
// 	var newCar models.Car

// 	if err := ctx.ShouldBindJSON(&newCar); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}
// 	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
// 	CarDatas = append(CarDatas, newCar)

// 	ctx.JSON(http.StatusCreated, gin.H{
// 		"car": newCar,
// 	})
// }

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var updatedCar models.Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			CarDatas[i] = updatedCar
			CarDatas[i].CarID = carID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully updated", carID),
	})
}

func GetAllCar(ctx *gin.Context) {
	condition := false
	// var carData Car

	if len(CarDatas) > 0 {
		condition = true
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("kaga ada mobilnya"),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": CarDatas,
	})
}

func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carData models.Car

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = models.Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully feleted", carID),
	})
}
