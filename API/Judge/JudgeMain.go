package Judge

import (
	"fmt"
	"os/exec"
)

/*==================================
# todo
- 帰ってきたJudgeデータをパースする
- パースしたデータをMainに返す
==================================*/

func Main(data map[string]string) {
	file := "Judge/WorkSpace/Main." + data["Lang"]
	answerData := data["AnswerData"]

	if err := MakeWorkSpace(); err == nil {
		if err := MakeAnswerFile(file); err == nil {
			if err := CopyJudgeFile(data["QuestionID"]); err == nil {
				if err := InputAnswer(answerData, file); err == nil {
					if err := RunJudge(data["Lang"]); err == nil {
						fmt.Println("DONE!")
					} else {
						DelWorkSpace()
					}
				} else {
					DelWorkSpace()
				}
			} else {
				DelWorkSpace()
			}
		} else {
			DelWorkSpace()
		}
	} else {
		DelWorkSpace()
	}
}

func RunCmd(lang string) (ret string) {
	switch lang {
	case "rb":
		ret = "ruby WorkSpace/Main.rb"
	case "py":
		ret = "python WorkSpace/Main.py"
	}
	return
}

func RunJudge(lang string) (err error) {
	if err = exec.Command("sh", "-c", "./Judge/docker_container_start.sh ", RunCmd(lang)).Run(); err != nil {
		fmt.Println("Runコマンドが失敗しました")
	}
	return
}

func CopyJudgeFile(questionID string) (err error) {
	cmd := "cp Judge/Questions/" + questionID + "/input.txt " + "Judge/WorkSpace"
	if err = exec.Command("sh", "-c", cmd).Run(); err != nil {
		fmt.Println("inputFileのコピーが失敗しました")
	}
	return
}

func InputAnswer(answerData, file string) (err error) {
	cmd := "echo " + answerData + ">" + file
	if err = exec.Command("sh", "-c", cmd).Run(); err != nil {
		fmt.Println("Answerの入力が失敗しました")
	}
	return
}

func DelAnswerFile(file string) (err error) {
	if err = exec.Command("rm", file).Run(); err != nil {
		fmt.Println("AnswerFileの削除に失敗しました")
	}
	return
}

func MakeAnswerFile(file string) (err error) {
	if err = exec.Command("touch", file).Run(); err != nil {
		fmt.Println("AnswerFileの作成が失敗しました")
	}
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
