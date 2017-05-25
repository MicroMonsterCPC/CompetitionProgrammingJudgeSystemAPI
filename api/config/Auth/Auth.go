package Auth

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Basic() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (error, bool) {
		if username == "joe" && password == "secret" {
			return nil, true
		}
		return nil, false
	})
}
