package funcs

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type record struct {
	Date      string  `json:"date"`
	Timestamp int64   `json:"timestamp"`
	Money     float64 `json:"money"`
	Category  string  `json:"category"`
	Type      string  `json:"type"`
	Desc      string  `json:"desc"`
}
// 获取偏移的时间
func GetOffsetTime(offset int) (string, int64) {
	t := time.Now().AddDate(0, 0, offset)
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day()), t.Unix()
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

func readDataToRecordSlice() ([]*record, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "打开文件错误")
	}
	scanner := bufio.NewScanner(file)
	jsonSlice := make([]*record, 0)
	for scanner.Scan() {
		txt := []byte(scanner.Text());
		i := 0
		l := len(txt)
		record := &record{}
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
					return nil, errors.Wrap(err, "解析 timestamp 出错")
				}
				record.Timestamp = timestamp
				break
			case 2:
				money, err := strconv.ParseFloat(strings.Trim(string(r), " "), 64)
				if err != nil {
					return nil, errors.Wrap(err, "解析 money 出错")
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
		jsonSlice = append(jsonSlice, record)
	}
	return jsonSlice, nil
}

// 记录要输出的数据 行号和行内容
type lineValue struct {
	value      []byte
	lineNumber int
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

// 获取文件的行数
func getFileLines(reader io.Reader) int {
	bioReade := bufio.NewReader(reader)
	count := 0
	for {
		_, err := bioReade.ReadString('\n')
		if err != nil {
			break
		}
		count++
	}
	return count
}
