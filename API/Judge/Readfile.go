package Judge

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
		fmt.Println("ファイルが見つかりませんでした")
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
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
