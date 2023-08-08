package handler

import (
	"github.com/labstack/echo/v4"
	"main/domain"
	"main/service"
	"net/http"
	"strconv"
)

func GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	res, err := service.GetUser(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func GetUser(c echo.Context) error {
	id := c.Get("id").(int)
	res, err := service.GetUser(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func CreateUser(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	res := service.CreateUser(*user)
	return c.JSON(http.StatusCreated, res)
}

func Login(c echo.Context) error {
	login := new(domain.Login)
	if err := c.Bind(login); err != nil {
		return err
	}

	isVaild, err := service.Login(*login)
	if err != nil {
		return err
	}

	if !isVaild {
		return c.NoContent(http.StatusUnauthorized)
	}

	jwt, err := service.NewJWT(login.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, jwt)
}
