package store

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/user"
)

type Record struct {
	Date      string  `json:"date"`
	Timestamp int64   `json:"timestamp"`
	Money     float64 `json:"money"`
	Category  string  `json:"category"`
	Type      string  `json:"type"`
	Desc      string  `json:"desc"`
}
type Species string

const (
	fileBackupPrefix         = "bookkeeping"
	INCOME           Species = "+"
	EXPENSE          Species = "-"
)

var (
	FileDirPath    = "/bookkeeping"
	FilePath       = "/bookkeeping/bookkeeping.txt"
	OutputFilePath       = "/bookkeeping/bookkeeping.json"
	FileBackupPath = "/bookkeeping/backup"
)

func init() {
	userInfo, _ := user.Current()
	FileDirPath = fmt.Sprintf("%s%s", userInfo.HomeDir, FileDirPath)
	FilePath = fmt.Sprintf("%s%s", userInfo.HomeDir, FilePath)
	FileBackupPath = fmt.Sprintf("%s%s", userInfo.HomeDir, FileBackupPath)
	OutputFilePath = fmt.Sprintf("%s%s", userInfo.HomeDir, OutputFilePath)
}

func Append(data *Record) error {
	str, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "序列化 json error")
	}
	f, err := os.OpenFile(FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("打开文件出错 file path: %s", FilePath))
	}
	defer f.Close()
	n, err := f.Write([]byte(str))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("写入文件出错 file path: %s", FilePath))
	}
	fmtString := "%-15s %-15s %-15s %-10s %-10s %-15s\n"
	fmt.Printf(fmtString, "date", "timestamp", "money", "category", "type", "description")
	fmtString = "%-15s %-15d %-15.2f %-10s %-10s %-15s\n"
	fmt.Printf(fmtString, data.Date, data.Timestamp, data.Money, data.Category, data.Type, data.Desc)
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
		str, err := json.Marshal(r)
		if err != nil {
			return errors.Wrap(err, "序列化 json error")
		}
		l, err := f.Write([]byte(str))
		l, err = f.WriteString("\n")
		length += l
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("写入文件出错 file path: %s", FilePath))
		}
	}
	fmt.Printf("ok, 共%d个字节\n", length)
	return nil
}
