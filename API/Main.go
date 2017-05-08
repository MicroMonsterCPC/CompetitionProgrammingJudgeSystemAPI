package main

import (
	"./Judge"
	"github.com/labstack/echo"
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
	e.GET("/", homePage)
	e.POST("/answer-data", answerData)
	e.Logger.Fatal(e.Start(":1323"))
}

func answerData(c echo.Context) error {
	data := new(answerDataJson)
	if err := c.Bind(data); err != nil {
		return err
	}
	userResult := map[string]string{
		"QuestionID": data.QuestionID,
		"AnswerData": data.AnswerData,
		"Lang":       data.Lang,
	}
	judgeResult := Judge.Main(userResult) //受け取った解答データをReadに投げる
	result := map[string]string{
		"Result": judgeResult[0],
	}
	return c.JSON(http.StatusOK, result)
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "This is a API")
}
