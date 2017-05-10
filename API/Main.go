package main

import (
	"./CreateQuestion"
	"./Judge"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type (
	answerDataJson struct {
		QuestionID string `json:"question_id"`
		AnswerData string `json:"answer_data"`
		Lang       string `json:"lang"`
	}
)

func main() {
	e := echo.New()

	//=======Middleware=======
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//=======Routers=======
	e.GET("/", homePage)
	e.POST("/answer-data", answerData)

	e.Logger.Fatal(e.Start(":1323"))
}

func answerData(c echo.Context) error {
	//POSTされてきたJSONデータをMAPに変換してJudgeに流す
	data := new(answerDataJson)
	if err := c.Bind(data); err != nil {
		return err
	}
	userResult := map[string]string{
		"QuestionID": data.QuestionID,
		"AnswerData": data.AnswerData,
		"Lang":       data.Lang,
	}
	var judgeResult []map[string]string = Judge.Main(userResult)
	//受け取ったUserの解答データをJudgeに投げる
	result := map[string][]map[string]string{
		"Result": judgeResult,
	}
	return c.JSON(http.StatusOK, result)
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "This is a API")
}
