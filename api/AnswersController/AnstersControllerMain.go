package AnswersController

import (
	"fmt"
	"io/ioutil"
	"os"
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
	content := []byte(AnswerData["AnswerData"])
	filename := "/echo-server/Judge/Questions/" + AnswerData["QuestionID"] + ".txt"
	fmt.Println(filename)
	if err := ioutil.WriteFile(filename, content, os.ModePerm); err != err {
		ret = false
		return
	}
	ret = true
	return
	// pwd, err := exec.Command("pwd").Output()
	// if err != nil {
	// 	ret = false
	// 	return
	// }
	// string(pwd)
}

func Update(AnswerData map[string]string) (ret bool) {
	return
}

func Delete(AnswerData map[string]string) (ret bool) {
	cmd := "rm ./Judge/Questions/" + AnswerData["QuestionID"] + ".txt"
	if err := exec.Command(cmd).Run; err != nil {
		ret = false
	}
	ret = true
	return
}
