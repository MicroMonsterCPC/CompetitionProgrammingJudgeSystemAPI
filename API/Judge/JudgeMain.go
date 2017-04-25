package Judge

import (
	"fmt"
	"os/exec"
)

/*==================================
# todo
- パースしたデータをDockerに流し込む
- 帰ってきたJudgeデータをパースする
- パースしたデータをMainに返す
==================================*/

func Main(data map[string]string) {
	file := "Judge/WorkSpace/Main." + data["Lang"]
	answerData := data["AnswerData"]

	if err := MakeWorkSpace(); err == nil {
		if err := MakeAnswerFile(file); err == nil {
			if err := InputAnswer(answerData, file); err == nil {
				if err := RunJudge(file, data["Lang"], data["QuestionID"]); err == nil {
					fmt.Println("DONE!")
					DelWorkSpace()
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

func RunCmd(file, lang string) (ret string) {
	switch lang {
	case "rb":
		ret = "ruby " + file
		fmt.Println(ret)
	case "py":
		ret = "python " + file
	}
	return
}

func RunJudge(file, lang, id string) (err error) {
	cmd := "Judge/run.sh " + RunCmd(file, lang) + " " + id
	fmt.Println(cmd)
	if err = exec.Command("sh", "-c", cmd).Run(); err != nil {
		fmt.Println("Runコマンドが失敗しました")
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
	return
}
func DelWorkSpace() (err error) {
	if err = exec.Command("rm", "-rf", "Judge/WorkSpace").Run(); err != nil {
		fmt.Println("WorkSpaceの削除に失敗しました")
	}
	return
}

func MakeWorkSpace() (err error) {
	if err = exec.Command("mkdir", "Judge/WorkSpace").Run(); err != nil {
		fmt.Println("WorkSpaceの作成が失敗しました")
	}
	return
}
