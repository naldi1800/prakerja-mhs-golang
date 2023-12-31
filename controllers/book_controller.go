package controllers

import (
	"net/http"
	"prakerja8/configs"
	"prakerja8/models"

	"github.com/labstack/echo/v4"
)


func AddBookController(c echo.Context) error {
	var requestBook models.Book
	c.Bind(&requestBook)

	// masukkan ke database
	result := configs.DB.Create(&requestBook)
	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false,
			Message: "Failed insert data books",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: requestBook,
	})

}

func GetDetailBookController(c echo.Context) error {
	
	// id, _ := strconv.Atoi(c.Param("id"))
	
	var book models.Book = models.Book{}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: book,
	})
	
}

func GetBooksController(c echo.Context) error {

	var books []models.Book
	
	result := configs.DB.Find(&books)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false,
			Message: "Failed insert data books",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: books,
	})
}
