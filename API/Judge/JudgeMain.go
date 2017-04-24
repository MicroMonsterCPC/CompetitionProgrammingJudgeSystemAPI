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
	file := "WorkSpace/Main." + data["Lang"]
	// questionFile := "Questions/" + data["QuestionID"] + "/input.txt"
	answerData := data["AnswerData"]

	if err := MakeWorkSpace(); err == nil {
		if err := MakeAnswerFile(file); err == nil {
			if err := InputAnswer(answerData, file); err == nil {
				fmt.Println("DONE!")
			} else {
				DelAnswerFile(file)
			}
		} else {
			DelAnswerFile(file)
		}
	} else {
		DelWorkSpace()
	}
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
	if err = exec.Command("rm", "-rf", "WorkSpace").Run(); err != nil {
		fmt.Println("WorkSpaceの削除に失敗しました")
	}
	return
}

func MakeWorkSpace() (err error) {
	if err = exec.Command("mkdir", "WorkSpace").Run(); err != nil {
		fmt.Println("WorkSpaceの作成が失敗しました")
	}
	return
}
