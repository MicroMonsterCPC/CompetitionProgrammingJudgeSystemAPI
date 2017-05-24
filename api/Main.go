package main

import (
	"./apps/AnswersController"
	"./apps/HomeController"
	"./apps/JudgementController"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	//=======Middleware=======
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//=======Routers=======
	e.GET("/", HomeController.Root)
	e.POST("/judgement-answer", JudgementController.Execution)
	// e.POST("/test-answer", testAnswers)

	e.POST("/answer", AnswersController.Create)
	// e.PUT("/answer", AnswersController.Update()
	e.DELETE("/answer", AnswersController.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
