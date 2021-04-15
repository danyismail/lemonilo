package handler

import (
	"fmt"
	"lemonilo/router/middleware"
	"lemonilo/utils"
	"net/http"

	// "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)

	v1.GET("/ping", h.Ping)

	auth := v1.Group("/auth")
	user := v1.Group("/user", jwtMiddleware)

	auth.POST("/register", h.CreateUser)
	auth.POST("/login", h.Login)

	user.GET("/list", h.ListUser)
	user.GET("/detail", h.DetailUser)
	user.POST("/update", h.UpdateUser)
	user.DELETE("/delete", h.DeleteUser)

}

func (h *Handler) HttpErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		var code = report.Code
		if report.Code > 88000 {
			code = http.StatusInternalServerError
		}
		rid := c.Response().Header().Get(echo.HeaderXRequestID)
		report.SetInternal(echo.NewHTTPError(0, "Request ID : "+rid))

		c.Logger().Error(report)

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required",
						err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email",
						err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param())
				}

				break
			}
		}

		c.JSON(code, report)
	}
}
