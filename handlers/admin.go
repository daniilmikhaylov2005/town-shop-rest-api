package handlers

import (
	"fmt"
	"github.com/daniilmikhaylov2005/town-shop-rest-api/middleware"
	"github.com/daniilmikhaylov2005/town-shop-rest-api/models"
	"github.com/daniilmikhaylov2005/town-shop-rest-api/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func Signup(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, response{
			Status: "Can't create user with this data",
		})
	}

	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Username or password can't be empty",
		})
	}

	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: "Can't hash password",
		})
	}

	user.Password = hashPassword
	user.Role = "user"

	id, err := repository.InsertUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: fmt.Sprintf("%d", id),
	})
}

func Signin(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Can't login with this data",
		})
	}

	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Username or password can't be empty",
		})
	}

	userDb, err := repository.FindUserByUsername(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}

	if !CheckPasswordHash(user.Password, userDb.Password) {
		return c.JSON(http.StatusBadRequest, response{
			Status: "Wrong password",
		})
	}

	user.Role = userDb.Role
	user.Name = userDb.Name

	token, err := CreateToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, responseToken{
		Token: token,
	})

}

func InsertGood(c echo.Context) error {
	claims, err := middleware.GetClaimsFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}
	if claims.Role != "admin" {
		return c.JSON(http.StatusForbidden, response{
			Status: "You don't have permission",
		})
	}

	var good models.Good

	if err := c.Bind(&good); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status: fmt.Sprintf("%v", err),
		})
	}
	if strings.TrimSpace(good.Name) == "" {
		return c.JSON(http.StatusBadRequest, response{
			Status: "You can't create good without name",
		})
	}
	if strings.TrimSpace(good.Image) == "" {
		good.Image = "empty"
	}
	if strings.TrimSpace(good.Description) == "" {
		return c.JSON(http.StatusBadRequest, response{
			Status: "You can't create good without description",
		})
	}
	if strings.TrimSpace(good.Category) == "" {
		return c.JSON(http.StatusBadRequest, response{
			Status: "You can't create good without category",
		})
	}

	id, err := repository.InsertGood(good)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: "Error while inserting good",
		})
	}

	good.ID = id

	return c.JSON(http.StatusCreated, good)
}
func UpdateGood(c echo.Context) error {
	claims, err := middleware.GetClaimsFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}
	if claims.Role != "admin" {
		return c.JSON(http.StatusForbidden, response{
			Status: "You don't have permission",
		})
	}
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status: err.Error(),
		})
	}

	var good models.Good
	if err := c.Bind(&good); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status: err.Error(),
		})
	}

	goodFromDb, err := repository.GetGoodById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}

	if strings.TrimSpace(good.Name) == "" {
		good.Name = goodFromDb.Name
	}
	if strings.TrimSpace(good.Description) == "" {
		good.Description = goodFromDb.Description
	}
	if strings.TrimSpace(good.Category) == "" {
		good.Category = goodFromDb.Category
	}

	goodId, err := repository.UpdateGood(good, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}

	good.ID = goodId

	return c.JSON(http.StatusOK, good)
}
func DeleteGood(c echo.Context) error {
	claims, err := middleware.GetClaimsFromJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}
	if claims.Role != "admin" {
		return c.JSON(http.StatusForbidden, response{
			Status: "You don't have permission",
		})
	}
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status: err.Error(),
		})
	}

	deletedId, err := repository.DeleteGood(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response{
		Status: fmt.Sprintf("good with id %d deleted", deletedId),
	})
}
