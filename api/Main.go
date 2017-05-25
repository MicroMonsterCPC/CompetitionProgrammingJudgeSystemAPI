package main

import (
	"./config/Router"
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

	//=======Router=======
	e := Router.NewRouter()

	//=======Middleware=======
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//=======TemplateEngine=======
	e.Renderer = t

	//=======Routers=======
	e.Logger.Fatal(e.Start(":1323"))
}
