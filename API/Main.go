package main

import (
	"./Judge"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", homePage)
	e.POST("/answer-data", answerData)
	e.Logger.Fatal(e.Start(":1323"))
}

func answerData(c echo.Context) error {
	Judge.Read() //受け取った解答データをReadに投げる
	result := map[string]string{
		"foo":  "bar",
		"hoge": "fuga",
	}
	return c.JSON(http.StatusOK, result)
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "This is a API")
}
