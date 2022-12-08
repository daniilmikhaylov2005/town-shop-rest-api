package handlers

import (
	"fmt"
	"github.com/daniilmikhaylov2005/town-shop-rest-api/models"
	"github.com/daniilmikhaylov2005/town-shop-rest-api/repository"
	"github.com/labstack/echo/v4"
	"net/http"
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
