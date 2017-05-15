package Judge

/*==================================
#役割
MAPにされた解答データを、
Dockerに流し込んで
実行結果ファイルを作りReadfileに投げる
==================================*/

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func Main(data map[string]string) (ret []map[string]string) {
	file := "Judge/WorkSpace/Main." + data["Lang"]
	answerData := data["AnswerData"]

	for {
		if err := MakeWorkSpace(); err != nil {
			fmt.Println(err)
			break
		}
		if err := MakeAnswerFile(); err != nil {
			break
		}
		if err := CopyJudgeFile(data["QuestionID"]); err != nil {
			break
		}
		if err := InputAnswer(answerData, file); err != nil {
			break
		}
		if err := RunJudge(data["Lang"]); err != nil {
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
	runCmd := "./Judge/docker_container_start.sh " + cmd + " " + image
	if err = exec.Command(runCmd).Run(); err != nil {
		fmt.Println(err)
		fmt.Println("Runコマンドが失敗しました")
	}
	return
}

func CopyJudgeFile(questionID string) (err error) {
	cmd := "cp Judge/Questions/" + questionID + ".txt " + "Judge/WorkSpace/input.txt"
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
	copyCmd := "cp Judge/judge_run.sh Judge/WorkSpace"
	if err = exec.Command("sh", "-c", copyCmd).Run(); err != nil {
		fmt.Println("JudgeRun.shのコピーに失敗しました")
	}
	return
}
func DelWorkSpace() (err error) {
	if err = exec.Command("rm", "-rf", "Judge/WorkSpace").Run(); err != nil {
		fmt.Println("WorkSpaceの削除に失敗しました")
	} else {
		fmt.Println("WorkSpaceを削除しました")
	}
	return
}

func MakeWorkSpace() (err error) {
	if err = exec.Command("mkdir", "Judge/WorkSpace").Run(); err != nil {
		fmt.Println("WorkSpaceの作成が失敗しました")
	}
	return
}
