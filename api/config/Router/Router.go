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
	e.GET("/", HomeController.Root)
	e.POST("/judgement-answer", JudgementController.Execution, Auth.Basic())
	// e.POST("/test-answer", testAnswers)

	e.POST("/answer", AnswersController.Create, Auth.Basic())
	e.PUT("/answer", AnswersController.Update, Auth.Basic())
	e.DELETE("/answer", AnswersController.Delete, Auth.Basic())
	return e
}
