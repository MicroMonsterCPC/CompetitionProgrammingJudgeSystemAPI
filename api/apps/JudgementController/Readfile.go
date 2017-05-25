package JudgementController

/*==================================
#役割
Dockerコンテナ上で作られた問題の
実行結果ファイル(Data)
をで読み取ってMapとして返却する
==================================*/

import (
	"bufio"
	"fmt"
	"github.com/k0kubun/pp"
	"os"
	"strconv"
)

func Read() (ret []map[string]string) {
	file, err := os.Open(`Judge/WorkSpace/data`)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ファイルが見つかりませんでした")
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			fmt.Println(err)
			break
		}
		//ToDo: ここでData（,区切り）のデータをSplitしてMapに入れる
		value := map[string]string{
			"ID":     strconv.Itoa(i),
			"Result": sc.Text(),
		}
		ret = append(ret, value)
	}
	pp.Println(ret)
	return
}
