package AnswersController

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

type (
	WantoToCreateAnswerJson struct {
		QuestionID string `json:"id" xml:"id" form:"" query:"id"`
		AnswerData string `json:"answer" xml:"answer" form:"answer" query:"answer"`
	}
)

func buildJSON(c echo.Context) (ret map[string]string) {
	WantoToCreateAnswer := new(WantoToCreateAnswerJson)
	if err := c.Bind(WantoToCreateAnswer); err != nil {
		fmt.Println("[E]: Bind Error")
		pp.Println(err)
	}
	pp.Println(WantoToCreateAnswer)
	ret = map[string]string{
		"QuestionID": WantoToCreateAnswer.QuestionID,
		"AnswerData": WantoToCreateAnswer.AnswerData,
	}
	return
}

func resultParse(value bool) (ret map[string]bool) {
	ret = map[string]bool{"Result": value}
	return
}

/*======================
controllers
======================== */

func Create(c echo.Context) error {
	var result map[string]bool
	AnswerData := buildJSON(c)

	content := []byte(AnswerData["AnswerData"])
	filename := "/echo-server/Judge/Questions/" + AnswerData["QuestionID"] + ".txt"
	fmt.Println(filename)
	if err := ioutil.WriteFile(filename, content, os.ModePerm); err != err {
		result = resultParse(false)
	} else {
		result = resultParse(true)
	}
	return c.JSON(http.StatusOK, result)
}

// func Update(c echo.Context) {
// 	Create()
// }

func Delete(c echo.Context) error {
	var result map[string]bool
	AnswerData := buildJSON(c)

	cmd := "rm ./Judge/Questions/" + AnswerData["QuestionID"] + ".txt"
	if err := exec.Command(cmd).Run; err != nil {
		result = resultParse(false)
	} else {
		result = resultParse(true)
	}
	return c.JSON(http.StatusOK, result)
}
