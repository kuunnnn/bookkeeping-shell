package read_data

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func Append(data *Record) error {
	fmtString := "%-15s %-15d %-15.2f %-10s %-10s %-15s\n"
	str := fmt.Sprintf(fmtString, data.Date, data.Timestamp, data.Money, data.Category, data.Type, data.Desc)
	f, err := os.OpenFile(FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("打开文件出错 file path: %s", FilePath))
	}
	defer f.Close()
	n, err := f.Write([]byte(str))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("写入文件出错 file path: %s", FilePath))
	}
	fmtString = "%-15s %-15s %-15s %-10s %-10s %-15s\n"
	fmt.Printf(fmtString, "date", "timestamp", "money", "category", "type", "description")
	fmt.Println(str)
	fmt.Printf("ok, 写入%d个字节\n", n)
	return nil
}

func Rewrite(data []*Record) error {
	f, err := os.OpenFile(FilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("打开文件出错 file path: %s", FilePath))
	}
	defer f.Close()
	length := 0
	for _, r := range data {
		str := fmt.Sprintf("%-15s %-15d %-15.2f %-10s %-10s %-15s\n", r.Date, r.Timestamp, r.Money, r.Category, r.Type, r.Desc)
		l, err := f.Write([]byte(str))
		length += l
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("写入文件出错 file path: %s", FilePath))
		}
	}
	fmt.Printf("ok, 共%d个字节\n", length)
	return nil
}
