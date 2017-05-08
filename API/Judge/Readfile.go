package Judge

import (
	"bufio"
	"fmt"
	"github.com/k0kubun/pp"
	"os"
)

func Read() (ret []string) {
	file, err := os.Open(`Judge/WorkSpace/data`)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			break
		}
		t := sc.Text()
		ret = append(ret, t)
	}
	pp.Println(ret)
	return
}
