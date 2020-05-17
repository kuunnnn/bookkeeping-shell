package funcs

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

func ToJSON() error {
	jsonSlice, err := readDataToRecordSlice();
	if err != nil {
		return errors.WithStack(err)
	}
	output := []string{"[\n"}
	for i, _ := range jsonSlice {
		s, _ := json.Marshal(jsonSlice[i])
		output = append(output, string(s), "\n")
		if i != len(jsonSlice)-1 {
			output = append(output, ",\n")
		}
	}
	output = append(output, "]")
	err = ioutil.WriteFile(outputFilePath, []byte(strings.Join(output, "")), 0644)
	if err != nil {
		return errors.Wrap(err, "写入 json 出错")
	}
	fmt.Printf("The file has been generated to %s\n", outputFilePath)
	return nil
}
