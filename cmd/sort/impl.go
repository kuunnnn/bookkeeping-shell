package sort

import (
	read_data "bookkeeping-shell/read-data"
	"github.com/pkg/errors"
)

func Sort() error {
	jsonSlice, err := read_data.ReadDataToRecordSlice()
	if err != nil {
		return errors.WithStack(err)
	}
	sort(jsonSlice)
	return read_data.Rewrite(jsonSlice)
}

// 进行排序
func sort(array []*read_data.Record) {
	i := 0
	j := len(array) - 1
	quickSort(array, i, j)
}

// 快排
func quickSort(array []*read_data.Record, low, high int) {
	if high-low <= 0 {
		return
	}
	i := low
	j := high
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
