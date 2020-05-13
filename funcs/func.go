package funcs

import (
	"bufio"
	json2 "encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
)

type TYPE string

const (
	FileBackupPrefix      = "bookkeeping"
	INCOME           TYPE = "+"
	EXPENSE          TYPE = "-"
)

var (
	FileDirPath    = "/bookkeeping"
	FilePath       = "/bookkeeping/bookkeeping.txt"
	OutputFilePath = "/bookkeeping/bookkeeping.json"
	FileBackupPath = "/bookkeeping/backup"
)

func init() {
	userInfo, _ := user.Current()
	FileDirPath = fmt.Sprintf("%s%s", userInfo.HomeDir, FileDirPath)
	FilePath = fmt.Sprintf("%s%s", userInfo.HomeDir, FilePath)
	FileBackupPath = fmt.Sprintf("%s%s", userInfo.HomeDir, FileBackupPath)
	FileBackupPath = fmt.Sprintf("%s%s", userInfo.HomeDir, OutputFilePath)
}

// 获取偏移的时间
func GetOffsetTime(offset int) (string, int64) {
	t := time.Now().AddDate(0, 0, offset)
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day()), t.Unix()
}

// 备份文件
// 备份后的文件后跟备份当天的日期
// 一天只写入一个文件
func Backup() error {
	srcFile, err := os.Open(FilePath)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("打开备份源文件错误 file path: %s", FilePath))
	}
	defer srcFile.Close()
	d, _ := GetOffsetTime(0)
	dst := fmt.Sprintf("%s/%s-%s.txt", FileBackupPath, FileBackupPrefix, d)
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
func Add(money float64, kind string, desc string, day int, typ TYPE) error {
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

// 切片是否不是空的
func byteSliceIsNotEmpty(byt []byte) bool {
	if byt == nil {
		return false
	}
	isNotEmpty := false
	for i := 0; i < len(byt); i++ {
		if byt[i] != 0 {
			isNotEmpty = true
			break
		}
	}
	return isNotEmpty
}

// 反转 []byte
func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

//反转 []*lineValue
func reverse2(s []*lineValue) []*lineValue {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

type lineValue struct {
	value      []byte
	lineNumber int
}

// 获取文件的行数
func getFileLines(reader io.Reader) int {
	breader := bufio.NewReader(reader)
	count := 0
	for {
		_, err := breader.ReadString('\n')
		if err != nil {
			break
		}
		count++
	}
	return count
}

// 输出最近[limit]条数据 并打印行号
func Cat(limit uint, offsetStart uint) error {
	file, err := os.Open(FilePath)
	if err != nil {
		return errors.Wrap(err, "打开文件错误")
	}
	defer file.Close()
	// 行号 倒数的
	lineNumber := getFileLines(file)
	// 定位到最后读到的第一个字节就是 \n 所以总行数需要加 1
	lineNumber++
	if offsetStart >= uint(lineNumber) {
		fmt.Printf("偏移量过大 all line %d offset %d\n", lineNumber, offsetStart)
		return nil
	}
	// 定位到文件最后面
	offset, err := file.Seek(0, os.SEEK_END)
	if err != nil {
		return errors.Wrap(err, "Seek err")
	}
	// 一次读一个字节
	buf := make([]byte, 1)
	// 缓冲一行数据大概的容量 先设置为 100 不够会自动扩容的
	lineBuf := make([]byte, 0)
	// 已经读了几行
	var line uint = 0
	// 存储读出的 byte
	lineSlice := make([]*lineValue, 0)
	// 一直读到 limit 行或者 读完
	i := offset - 1
	for ; i >= 0; i-- {
		_, err = file.ReadAt(buf, i)
		// 从后面往前面读不会遇到 io.EOF 错误
		if err != nil {
			return errors.Wrap(err, "ReadAt err")
		}
		// 如果没有读到上一行
		if buf[0] != 10 {
			lineBuf = append(lineBuf, buf[0])
			continue
		}
		// 跳过空行
		// 定位到最后读到的第一个字节就是 \n 也就是说一定会有一个空行
		if byteSliceIsNotEmpty(lineBuf) {
			if offsetStart == 0 {
				line = line + 1
				// 保存这一行数据
				lineSlice = append(lineSlice, &lineValue{
					lineNumber: lineNumber,
					value:      reverse(lineBuf),
				})
			} else {
				offsetStart--
			}
			// 清空 lineBuf
			lineBuf = make([]byte, 0)
			// 已经读够行数了
			if line == limit {
				break
			}
		}
		lineNumber--
	}
	// 如果 i 是 -1 表示 没有足够的非空行
	if i == -1 {
		lineSlice = append(lineSlice, &lineValue{
			lineNumber: lineNumber,
			value:      reverse(lineBuf),
		})
	}
	template := "%-15s %-15s %-15s %-15s %-10s %-10s %-15s\n"
	fmt.Printf(template, "Lines", "date", "timestamp", "money", "category", "type", "description")
	for _, v := range reverse2(lineSlice) {
		if byteSliceIsNotEmpty(v.value) {
			fmt.Printf("%-15d %-15s\n", v.lineNumber, string(v.value))
		}
	}
	return nil
}

// 删除第[line]行数据
func Delete(line uint) error {
	if line == 0 {
		return nil
	}
	file, err := os.Open(FilePath)
	if err != nil {
		return errors.Wrap(err, "打开文件错误")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	buf := make([]byte, 0)
	i := 1
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
	err = ioutil.WriteFile(FilePath, buf, 0644)
	if err != nil {
		return errors.Wrap(err, "Write error")
	}
	fmt.Printf("Deleted: %d ok!\n", line)
	return nil
}

type Record struct {
	Date      string  `json:"date"`
	Timestamp int64   `json:"timestamp"`
	Money     float64 `json:"money"`
	Category  string  `json:"category"`
	Type      string  `json:"type"`
	Desc      string  `json:"desc"`
}

func ToJSON() error {
	file, err := os.Open(FilePath)
	if err != nil {
		return errors.Wrap(err, "打开文件错误")
	}
	scanner := bufio.NewScanner(file)
	json := make([]*Record, 0)
	for scanner.Scan() {
		txt := []byte(scanner.Text());
		i := 0
		l := len(txt)
		record := &Record{}
		for j := 0; i < l; j++ {
			r, i1 := readString(txt, i, l)
			i = i1
			switch j {
			case 0:
				record.Date = string(r)
				break
			case 1:
				timestamp, err := strconv.ParseInt(strings.Trim(string(r), " "), 10, 64)
				if err != nil {
					return errors.Wrap(err, "解析 timestamp 出错")
				}
				record.Timestamp = timestamp
				break
			case 2:
				money, err := strconv.ParseFloat(strings.Trim(string(r), " "), 64)
				if err != nil {
					return errors.Wrap(err, "解析 money 出错")
				}
				record.Money = money
				break
			case 3:
				record.Category = string(r)
				break
			case 4:
				record.Type = string(r)
				break
			case 5:
				record.Desc = string(r)
				break
			}
		}
		json = append(json, record)
	}
	output := []string{"[\n"}
	for i, _ := range json {
		s, _ := json2.Marshal(json[i])
		output = append(output, string(s), "\n")
		if i != len(json)-1 {
			output = append(output, ",\n")
		}
	}
	output = append(output, "]")
	err = ioutil.WriteFile(OutputFilePath, []byte(strings.Join(output, "")), 0644)
	if err != nil {
		return errors.Wrap(err, "写入 json 出错")
	}
	fmt.Printf("The file has been generated to %s\n", OutputFilePath)
	return nil
}

// 读取字节直到 空格 并跳过连着的空格
func readString(list []byte, i int, l int) ([]byte, int) {
	result := make([]byte, 0)
	for ; i < l; i++ {
		if list[i] == 32 {
			i++
			for i < l {
				if list[i] != 32 {
					return result, i
				}
				i++
			}
		} else {
			result = append(result, list[i])
		}
	}
	return result, i
}
