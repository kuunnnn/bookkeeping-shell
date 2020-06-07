package funcs

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

// 删除第[line]行数据
func Delete(line uint) error {
	if line == 0 {
		return nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		return errors.Wrap(err, "打开文件错误")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	buf := make([]byte, 0)
	i := 1
	// 按照行读取并跳过给定的行号, 然后重新写入全部数据
	// TODO 优化写入是否可以只写部分
	for {
		b, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		if i != int(line) {
			buf = append(buf, b...)
		}
		i++
	}
	if i < int(line) {
		return errors.New(fmt.Sprintf("The line was not found [%d] \n", line))
	}
	err = ioutil.WriteFile(filePath, buf, 0644)
	if err != nil {
		return errors.Wrap(err, "Write error")
	}
	fmt.Printf("Deleted: %d ok!\n", line)
	return nil
}
