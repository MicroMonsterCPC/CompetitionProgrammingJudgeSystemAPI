package main

import (
	"./AnswersController"
	"./Judge"
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type (
	answerDataJson struct {
		QuestionID string `json:"id" xml:"id" form:"name" query:"name"`
		AnswerData string `json:"code" xml:"code" form:"code" query:"name"`
		Lang       string `json:"lang" xml:"lang" form:"lang" qyery:"lang"`
	}
	//Client側では使わない
	WantoToCreateAnswerJson struct {
		Action     string `json:"action" xml:"action" form:"action" query:"action"`
		QuestionID string `json:"id" xml:"id" form:"" query:"id"`
		AnswerData string `json:"answer" xml:"answer" form:"answer" query:"answer"`
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
	e.POST("/create-answer", cudAnswer)

	e.Logger.Fatal(e.Start(":1323"))
}

func cudAnswer(c echo.Context) error {
	WantoToCreateAnswer := new(WantoToCreateAnswerJson)
	if err := c.Bind(WantoToCreateAnswer); err != nil {
		pp.Println(err)
		return err
	}
	pp.Println(WantoToCreateAnswer)
	getData := map[string]string{
		"Action":     WantoToCreateAnswer.Action,
		"QuestionID": WantoToCreateAnswer.QuestionID,
		"AnswerData": WantoToCreateAnswer.AnswerData,
	}
	result := AnswersController.Main(getData) // map[string]bool
	return c.JSON(http.StatusOK, result)
}

func answerData(c echo.Context) error {
	//POSTされてきたJSONデータをMAPに変換してJudgeに流す
	data := new(answerDataJson)
	if err := c.Bind(data); err != nil {
		fmt.Println("[E]: Bind Error")
		fmt.Println(err)
		return err
	}
	pp.Println(data)
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
	return c.String(http.StatusOK, "This is a CompetitionProgrammingJudgeSystem the API")
}
