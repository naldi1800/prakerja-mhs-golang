package controllers

import (
	"net/http"
	"prakerja8/configs"
	"prakerja8/models"

	"github.com/labstack/echo/v4"
)

func AddMhsController(c echo.Context) error {
	var requestMhs models.Mhs
	c.Bind(&requestMhs)

	// masukkan ke database
	result := configs.DB.Create(&requestMhs)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed insert data mhs",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil",
		Data:    requestMhs,
	})
}

// func GetDetailMhsController(c echo.Context) error {
// 	// id, _ := strconv.Atoi(c.Param("id"))
// 	var mhs models.Mhs = models.Mhs{}
// 	return c.JSON(http.StatusOK, models.BaseResponse{
// 		Status:  true,
// 		Message: "Berhasil",
// 		Data:    mhs,
// 	})
// }

func GetAllMhsController(c echo.Context) error {

	var mhs []models.Mhs

	result := configs.DB.Find(&mhs)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed insert data mhs",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil",
		Data:    mhs,
	})
}
