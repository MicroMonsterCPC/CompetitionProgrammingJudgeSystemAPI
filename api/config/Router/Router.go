package Router

import (
	"../../apps/AnswersController"
	"../../apps/HomeController"
	"../../apps/JudgementController"
	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", HomeController.Root)
	e.POST("/judgement-answer", JudgementController.Execution)
	// e.POST("/test-answer", testAnswers)

	e.POST("/answer", AnswersController.Create)
	// e.PUT("/answer", AnswersController.Update()
	e.DELETE("/answer", AnswersController.Delete)
	return e
}
