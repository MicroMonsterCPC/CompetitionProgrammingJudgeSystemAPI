package Router

import (
	"../../apps/AnswersController"
	"../../apps/HomeController"
	"../../apps/JudgementController"
	"../../config/Auth"
	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", HomeController.Root, Auth.Basic())
	e.POST("/judgement-answer", JudgementController.Execution)
	// e.POST("/test-answer", testAnswers)

	e.POST("/answer", AnswersController.Create)
	// e.PUT("/answer", AnswersController.Update()
	e.DELETE("/answer", AnswersController.Delete)
	return e
}
