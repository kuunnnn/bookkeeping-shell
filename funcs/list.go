package funcs

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

// 输出最近[limit]条数据 并打印行号
func Cat(limit uint, offsetStart uint) error {
	file, err := os.Open(filePath)
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
