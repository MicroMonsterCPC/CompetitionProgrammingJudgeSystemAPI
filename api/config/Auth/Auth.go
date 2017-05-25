package Auth

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func Basic() echo.MiddlewareFunc {
	if "" == os.Getenv("AUTH_USERNAME") && "" == os.Getenv("AUTH_PASSWORD") {
		fmt.Println("The environment variable is not set")
	}
	return middleware.BasicAuth(func(username, password string, c echo.Context) (error, bool) {
		if username == os.Getenv("AUTH_USERNAME") && password == os.Getenv("AUTH_PASSWORD") {
			return nil, true
		}
		return nil, false
	})
}
