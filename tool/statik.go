package tool

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"os"
)

type stackObject struct {
	dirs     []os.FileInfo
	basePath string
}

// 生成 view 文件
// 将 statik 打包的静态资源生成到 指定的本地目录下
func GenDistFile(localBaseDir string) error {
	statikFS, err := fs.New()
	if err != nil {
		return errors.Wrap(err, "statik New Error")
	}
	r, err := statikFS.Open("/")
	if err != nil {
		return errors.Wrap(err, "open statik root dir Error")
	}
	defer r.Close()
	rootDir, err := r.Readdir(20)
	if err != nil {
		return errors.Wrap(err, "read statik root dir error")
	}
	localBaseDir = fmt.Sprintf("%s/view", localBaseDir)
	err = CreateDir(localBaseDir)
	if err != nil {
		return errors.Wrap(err, "create local dir error")
	}
	var fileInfoStack []*stackObject
	fileInfoStack = append(fileInfoStack, &stackObject{
		dirs:     rootDir,
		basePath: "",
	})
	for len(fileInfoStack) != 0 {
		object := fileInfoStack[0]
		fileInfoStack = fileInfoStack[1:]
		for _, file := range object.dirs {
			nextDirPath := fmt.Sprintf("%s/%s", object.basePath, file.Name())
			fileData, err := statikFS.Open(nextDirPath)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("open file[ %s ] error", nextDirPath))
			}
			if file.IsDir() {
				localNextDir := fmt.Sprintf("%s%s", localBaseDir, nextDirPath)
				err := CreateDir(localNextDir)
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("gen local dir[ %s ] error ", localNextDir))
				}
				dir, err := fileData.Readdir(100)
				fileData.Close()
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("read statik dir[ %s ] error", localNextDir))
				}
				fileInfoStack = append(fileInfoStack, &stackObject{
					basePath: nextDirPath,
					dirs:     dir,
				})
				continue
			}
			fileAllBytes, err := ioutil.ReadAll(fileData)
			fileData.Close()
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("read statik file[ %s ] bytes error", nextDirPath))
			}
			localFilePath := fmt.Sprintf("%s%s", localBaseDir, nextDirPath)
			err = ioutil.WriteFile(localFilePath, fileAllBytes, 0666)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("write file[ %s ] error", localFilePath))
			}
		}
	}
	return nil
}
