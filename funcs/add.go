package funcs

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

// 新增一条记录
func Add(money float64, kind string, desc string, day int, typ Species) error {
	d, t := GetOffsetTime(day)
	str := fmt.Sprintf("%-15s %-15d %-15.2f %-10s %-10s %-15s\n", d, t, money, kind, typ, desc)
	f, err := os.OpenFile(FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("打开文件出错 file path: %s", FilePath))
	}
	defer f.Close()
	n, err := f.Write([]byte(str))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("写入文件出错 file path: %s", FilePath))
	}
	fmt.Printf("%-15s %-15s %-15s %-10s %-10s %-15s\n", "date", "timestamp", "money", "category", "type", "description")
	fmt.Println(str)
	fmt.Printf("ok, 写入%d个字节\n", n)
	return nil
}
