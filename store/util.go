package store

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"time"
)

// 获取偏移的时间
func GetOffsetTime(offset int) (string, int64) {
	t := time.Now().AddDate(0, 0, offset)
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day()), t.Unix()
}

// 获取文件的行数
func GetFileLines(reader io.Reader) int {
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

func ReadDataToRecordSlice() ([]*Record, error) {
	file, err := os.Open(FilePath)
	if err != nil {
		return nil, errors.Wrap(err, "打开文件错误")
	}
	scanner := bufio.NewScanner(file)
	jsonSlice := make([]*Record, 0)
	for scanner.Scan() {
		txt := []byte(scanner.Text());
		record := &Record{}
		err := json.Unmarshal(txt, &record)
		if err != nil {
			return nil, errors.Wrap(err, "序列化 json 字符串 error")
		}
		jsonSlice = append(jsonSlice, record)
	}
	return jsonSlice, nil
}

// 记录要输出的数据 行号和行内容
type LineValue struct {
	Value      []byte
	LineNumber int
}

// 切片是否不是空的
func ByteSliceIsNotEmpty(byt []byte) bool {
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
func Reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

//反转 []*lineValue
func Reverse2(s []*LineValue) []*LineValue {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
