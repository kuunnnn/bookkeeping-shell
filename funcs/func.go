package funcs

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"os/user"
)

type Species string

const (
	fileBackupPrefix         = "bookkeeping"
	INCOME           Species = "+"
	EXPENSE          Species = "-"
)

var (
	fileDirPath    = "/bookkeeping"
	filePath       = "/bookkeeping/bookkeeping.txt"
	outputFilePath = "/bookkeeping/bookkeeping.json"
	FileBackupPath = "/bookkeeping/backup"
)

func init() {
	userInfo, _ := user.Current()
	fileDirPath = fmt.Sprintf("%s%s", userInfo.HomeDir, fileDirPath)
	filePath = fmt.Sprintf("%s%s", userInfo.HomeDir, filePath)
	FileBackupPath = fmt.Sprintf("%s%s", userInfo.HomeDir, FileBackupPath)
	outputFilePath = fmt.Sprintf("%s%s", userInfo.HomeDir, outputFilePath)
}

// 备份文件
// 备份后的文件后跟备份当天的日期
// 一天只写入一个文件
func Backup() error {
	srcFile, err := os.Open(filePath)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("打开备份源文件错误 file path: %s", filePath))
	}
	defer srcFile.Close()
	d, _ := GetOffsetTime(0)
	dst := fmt.Sprintf("%s/%s-%s.txt", FileBackupPath, fileBackupPrefix, d)
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("新建备份文件错误 file path: %s", dst))
	}
	defer dstFile.Close()
	n, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return errors.Wrap(err, "复制文件出错")
	}
	fmt.Printf("备份完成,写入%d个字节\n", n)
	return nil
}

// 新增一条记录
func Add(money float64, kind string, desc string, day int, typ Species) error {
	d, t := GetOffsetTime(day)
	str := fmt.Sprintf("%-15s %-15d %-15.2f %-10s %-10s %-15s\n", d, t, money, kind, typ, desc)
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("打开文件出错 file path: %s", filePath))
	}
	defer f.Close()
	n, err := f.Write([]byte(str))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("写入文件出错 file path: %s", filePath))
	}
	fmt.Printf("%-15s %-15s %-15s %-10s %-10s %-15s\n", "date", "timestamp", "money", "category", "type", "description")
	fmt.Println(str)
	fmt.Printf("ok, 写入%d个字节\n", n)
	return nil
}

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
