package JudgementController

/*==================================
#役割
MAPにされた解答データを、
Dockerに流し込んで
実行結果ファイルを作りReadfileに投げる
==================================*/

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
	answerDataJson struct {
		QuestionID string `json:"id" xml:"id" form:"name" query:"name"`
		AnswerData string `json:"code" xml:"code" form:"code" query:"name"`
		Lang       string `json:"lang" xml:"lang" form:"lang" qyery:"lang"`
	}
)

func Execution(c echo.Context) error {
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
	var judgeResult []map[string]string = Main(userResult)
	//受け取ったUserの解答データをJudgeに投げる
	result := map[string][]map[string]string{
		"Result": judgeResult,
	}
	return c.JSON(http.StatusOK, result)
}

func Main(data map[string]string) (ret []map[string]string) {
	file := "apps/JudgementController/WorkSpace/Main." + data["Lang"]
	answerData := data["AnswerData"]

	for {
		if err := MakeWorkSpace(); err != nil {
			fmt.Println(err)
			break
		}
		if err := MakeAnswerFile(); err != nil {
			fmt.Println(err)
			break
		}
		if err := CopyJudgeFile(data["QuestionID"]); err != nil {
			fmt.Println(err)
			break
		}
		if err := InputAnswer(answerData, file); err != nil {
			fmt.Println(err)
			break
		}
		if err := RunJudge(data["Lang"]); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("done")
		ret = Read() //retにReadから受け取った判定結果のMapを受け取る
	}
	DelWorkSpace()
	return
}

func RunCmd(lang string) (cmd, image string) {
	switch lang {
	case "rb":
		cmd = "ruby WorkSpace/Main.rb"
		image = "ruby:2.3.4-alpine"
	case "py":
		cmd = "python WorkSpace/Main.py"
		image = "python:2-alpine"
	}
	return
}

func RunJudge(lang string) (err error) {
	cmd, image := RunCmd(lang)
	runCmd := ".apps/JudgementController/docker_container_start.sh " + cmd + " " + image
	if err = exec.Command(runCmd).Run(); err != nil {
		fmt.Println(err)
		fmt.Println("Runコマンドが失敗しました")
	}
	return
}

func CopyJudgeFile(questionID string) (err error) {
	cmd := "cp apps/JudgementController/Questions/" + questionID + ".txt " + "apps/JudgementController/WorkSpace/input.txt"
	if err = exec.Command("sh", "-c", cmd).Run(); err != nil {
		fmt.Println("inputFileのコピーが失敗しました")
	}
	return
}

func InputAnswer(answerData, file string) (err error) {
	content := []byte(answerData)
	if err := ioutil.WriteFile(file, content, os.ModePerm); err != err {
		fmt.Println("AnswerFileの作成に失敗しました")
	}
	return
}

func MakeAnswerFile() (err error) {
	copyCmd := "cp apps/JudgementController/judge_run.sh apps/JudgementController/WorkSpace"
	if err = exec.Command("sh", "-c", copyCmd).Run(); err != nil {
		fmt.Println("JudgeRun.shのコピーに失敗しました")
	}
	return
}
func DelWorkSpace() (err error) {
	if err = exec.Command("rm", "-rf", "apps/JudgementController/WorkSpace").Run(); err != nil {
		fmt.Println("WorkSpaceの削除に失敗しました")
	} else {
		fmt.Println("WorkSpaceを削除しました")
	}
	return
}

func MakeWorkSpace() (err error) {
	if err = exec.Command("mkdir", "apps/JudgementController/WorkSpace").Run(); err != nil {
		fmt.Println("WorkSpaceの作成が失敗しました")
	}
	return
}
