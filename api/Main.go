package main

import (
	"./apps/AnswersController"
	"./apps/HomeController"
	"./apps/JudgementController"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()

	//=======Middleware=======
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//=======TemplateEngine=======
	e.Renderer = t

	//=======Routers=======
	e.GET("/", HomeController.Root)
	e.POST("/judgement-answer", JudgementController.Execution)
	// e.POST("/test-answer", testAnswers)

	e.POST("/answer", AnswersController.Create)
	// e.PUT("/answer", AnswersController.Update()
	e.DELETE("/answer", AnswersController.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
