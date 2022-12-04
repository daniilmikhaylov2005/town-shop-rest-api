package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type tempResponse struct {
	Category string `json: "category"`
	ID       int    `json: "id"`
}

func GetAllGoods(c echo.Context) error {
	category := c.Param("category")
	return c.JSON(http.StatusOK, tempResponse{
		Category: category,
	})
}

func GetGoodById(c echo.Context) error {
	category := c.Param("category")
	stringId := c.Param("id")
	intId, err := strconv.Atoi(stringId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status string `json:"status"`
		}{
			Status: "Error while convert id from url to string",
		})
	}
	return c.JSON(http.StatusOK, tempResponse{
		Category: category,
		ID:       intId,
	})
}
