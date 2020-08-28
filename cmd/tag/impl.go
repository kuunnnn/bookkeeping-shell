package tag

import (
	"bookkeeping-shell/store"
	"fmt"
	"github.com/pkg/errors"
)
func ArrayIncludes(array []string, name string) bool {
	for _, value := range array {
		if value == name {
			return true
		}
	}
	return false
}

func PrintCategoryList() error {
	data, err := store.ReadDataToRecordSlice()
	if err != nil {
		return errors.Wrap(err, "读取数据文件错误")
	}
	tagMap := make(map[string]float64)
	for _, record := range data {
		if money, ok := tagMap[record.Category]; ok {
			tagMap[record.Category] = money + record.Money
		} else {
			tagMap[record.Category] = record.Money
		}
	}
	for key, value := range tagMap {
		fmt.Printf("%s : %-.2f\n", key, value)
	}
	return nil
}
func ModifyCategoryName(fromName []string, targetName string) error {
	data, err := store.ReadDataToRecordSlice()
	if err != nil {
		return errors.Wrap(err, "读取数据文件错误")
	}
	for _, record := range data {
		if ArrayIncludes(fromName, record.Category) {
			record.Category = targetName
		}
	}
	return store.Rewrite(data)
}
