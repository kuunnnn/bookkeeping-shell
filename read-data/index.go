package read_data

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
	OutputFilePath = fmt.Sprintf("%s%s", userInfo.HomeDir, OutputFilePath)
}
