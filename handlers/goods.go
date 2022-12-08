package handlers

import (
	"github.com/daniilmikhaylov2005/town-shop-rest-api/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllGoods(c echo.Context) error {
	category := c.Param("category")

	goods, err := repository.GetAllGoods(category)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, goods)
}

func GetGoodById(c echo.Context) error {
	category := c.Param("category")
	stringId := c.Param("id")
	intId, err := strconv.Atoi(stringId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Error while get id from url",
		})
	}

	good, err := repository.GetGoodById(category, intId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, good)
}
