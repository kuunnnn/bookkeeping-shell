package funcs

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type BillType = int

const (
	MONTH BillType = 0
	YEAR  BillType = 1
	TAG   BillType = 2
)

func OutputBillData(typ BillType, species Species) error {
	recordSlice, err := readDataToRecordSlice()
	if err != nil {
		return errors.WithStack(err)
	}
	data := make(map[string]float64)
	for _, v := range recordSlice {
		key := ""
		switch typ {
		case YEAR:
			key = strconv.Itoa(int(time.Unix(v.Timestamp, 0).Year()))
			break
		case MONTH:
			key = strconv.Itoa(int(time.Unix(v.Timestamp, 0).Month()))
			break
		case TAG:
			key = v.Category
			break
		}
		if m, ok := data[key]; ok {
			data[key] = m + v.Money
		} else {
			data[key] = v.Money
		}
	}
	if len(data) == 0 {
		fmt.Println("暂无数据")
		return nil
	}
	label := ""
	switch typ {
	case YEAR:
		label = "年"
		break
	case MONTH:
		label = "月"
		break
	case TAG:
		label = "标签"
		break
	}
	for k, v := range data {
		switch species {
		case INCOME:
			fmt.Printf("%-4s %s 收入 %.2f\n", k, label, v)
			break
		case EXPENSE:
			fmt.Printf("%-4s %s 支出 %.2f\n", k, label, v)
			break
		default:

		}
	}
	return nil
}
