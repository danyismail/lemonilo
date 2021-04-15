package handler

import (
	"lemonilo/model"
	"lemonilo/response"
	"lemonilo/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Ping(c echo.Context) error {
	return c.JSON(200, "Pong")
}

func (h *Handler) CreateUser(c echo.Context) error {
	c.Logger().Info(".: call CreateUser Handler :.")

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    err.Error(),
		})
	}
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    err.Error(),
		})
	}

	user.Password = utils.HashAndSalt(user.Password)
	err := h.userStore.Create(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: true,
		Msg:    "successfully create user",
	})
}

func (h *Handler) ListUser(c echo.Context) error {
	c.Logger().Info(".: call ListUser Handler :.")
	result, err := h.userStore.List()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    err.Error(),
		})
	}
	var users []response.Users

	for _, val := range *result {
		var user = response.Users{
			ID:        val.ID,
			CreatedAt: utils.ParseDateToString(val.CreatedAt, utils.DateTimeFormat),
			UpdatedAt: utils.ParseDateToString(val.UpdatedAt, utils.DateTimeFormat),
			Email:     val.Email,
			Address:   val.Address,
		}
		users = append(users, user)
	}
	return c.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: true,
		Msg:    "successfully show list users",
		Data:   users,
	})
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.QueryParam("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    err.Error(),
		})
	}
	err = h.userStore.Delete(uint(i))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: true,
		Msg:    "successfully deleted user",
	})
}

func (h *Handler) DetailUser(c echo.Context) error {
	id := c.QueryParam("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    err.Error(),
		})
	}
	result, err := h.userStore.Detail(uint(i))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    err.Error(),
		})
	}

	if result == nil {
		return c.JSON(http.StatusOK, response.Response{
			Code:   http.StatusOK,
			Status: true,
			Msg:    "user not found",
		})
	}

	var user = response.Users{
		ID:        result.ID,
		CreatedAt: utils.ParseDateToString(result.CreatedAt, utils.DateTimeFormat),
		UpdatedAt: utils.ParseDateToString(result.CreatedAt, utils.DateTimeFormat),
		Email:     result.Email,
	}

	return c.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: true,
		Msg:    "successfully get detail user",
		Data:   user,
	})

}

func (h *Handler) UpdateUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(&user); err != nil {
		return err
	}

	user.Password = utils.HashAndSalt(user.Password)
	err := h.userStore.Update(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    "user not found",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: true,
		Msg:    "successfully update detail user",
	})
}

func (h *Handler) Login(c echo.Context) error {
	c.Logger().Info(".: call Login Handler :.")

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    err.Error(),
		})
	}
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    err.Error(),
		})
	}

	result, err := h.userStore.Login(user.Email)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    err.Error(),
		})
	}

	if result == nil {
		return c.JSON(http.StatusForbidden, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    "Authentication failed",
		})
	}

	if err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusForbidden, response.Response{
			Code:   http.StatusBadRequest,
			Status: false,
			Msg:    "Authentication failed",
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: true,
		Msg:    "Login Successfully",
		Data:   utils.GenerateJWT(result.Email),
	})
}
