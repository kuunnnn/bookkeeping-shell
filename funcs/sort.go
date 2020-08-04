package funcs

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)
func quickSort(array []*record, low, high int) {
	if high-low <=0{
		return
	}
	i:=low
	j:=high
	mid := array[i]
	for i < j {
		if array[j].Timestamp >= mid.Timestamp && i < j {
			j--
		}
		array[i] = array[j]
		if array[i].Timestamp <= mid.Timestamp && i < j {
			i++
		}
		array[j] = array[i]
	}
	array[i] = mid
	quickSort(array, low, i-1)
	quickSort(array, i+1, high)
}

func sort(array []*record) {
	i := 0
	j := len(array) - 1
	quickSort(array, i, j)
}

func Sort() error {
	jsonSlice, err := readDataToRecordSlice()
	if err != nil {
		return errors.WithStack(err)
	}
	sort(jsonSlice)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("打开文件出错 file path: %s", filePath))
	}
	defer f.Close()
	length:=0
	for _, r :=range jsonSlice  {
		str := fmt.Sprintf("%-15s %-15d %-15.2f %-10s %-10s %-15s\n", r.Date, r.Timestamp, r.Money, r.Category, r.Type, r.Desc)
		l, err := f.Write([]byte(str))
		length+=l
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("写入文件出错 file path: %s", filePath))
		}
	}
	fmt.Printf("ok, 共%d个字节\n", length)
	return nil
}
