package tool

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func IsFileExist(fileName string) (error, bool) {
	_, err := os.Stat(fileName)
	if err == nil {
		return nil, true
	}
	if os.IsNotExist(err) {
		return nil, false
	}
	return err, false
}

func CreateDir(p string) error {
	var err error
	err, isExist := IsFileExist(p)
	if err != nil {
		return errors.WithStack(err)
	}
	if !isExist {
		if err = os.MkdirAll(p, 0711); err != nil {
			return errors.Wrap(err,fmt.Sprintf("create dir error [ %s ]", p))
		}
	}
	return nil
}
