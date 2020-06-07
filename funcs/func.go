package funcs

import (
	"fmt"
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
