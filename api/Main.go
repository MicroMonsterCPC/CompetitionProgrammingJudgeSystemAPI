package main

import (
	"./AnswersController"
	"./Judge"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	//=======Middleware=======
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//=======Routers=======
	e.GET("/", homePage)
	e.POST("/judgement-answer", Judge.Execution)
	// e.POST("/test-answer", testAnswers)

	e.POST("/answer", AnswersController.Create)
	// e.PUT("/answer", AnswersController.Update()
	e.DELETE("/answer", AnswersController.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "This is a CompetitionProgrammingJudgeSystem the API")
}
