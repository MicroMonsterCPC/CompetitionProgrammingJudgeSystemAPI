package AnswersController

import (
	"fmt"
	"os/exec"
)

func Main(AnswerData map[string]string) (ret map[string]bool) {
	var result bool
	switch AnswerData["Action"] {
	case "create":
		result = Create(AnswerData)
	case "update":
		result = Update(AnswerData)
	case "delete":
		result = Delete(AnswerData)
	}
	ret = map[string]bool{
		"Result": result,
	}
	return
}

func Create(AnswerData map[string]string) (ret bool) {
	cmd := "echo " + AnswerData["AnswerData"] + " > ./Judge/Questions/" + AnswerData["QuestionID"] + ".txt"
	if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
		fmt.Println("AnswerFileの作成が失敗しました")
		ret = false
		return
	}
	ret = true
	return
}

func Update(AnswerData map[string]string) (ret bool) {
	return
}

func Delete(AnswerData map[string]string) (ret bool) {
	return
}
