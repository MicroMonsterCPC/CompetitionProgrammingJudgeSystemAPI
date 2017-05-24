package HomeController

import (
	"github.com/labstack/echo"
	"net/http"
)

func Root(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "world")
	// return c.String(http.StatusOK, "This is a CompetitionProgrammingJudgeSystem the API")
}
