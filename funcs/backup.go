package funcs

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
)

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

