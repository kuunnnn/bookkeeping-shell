package conv

import (
	read_data "bookkeeping-shell/read-data"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

func ToJSON() error {
	jsonSlice, err := read_data.ReadDataToRecordSlice()
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
	err = ioutil.WriteFile(read_data.OutputFilePath, []byte(strings.Join(output, "")), 0644)
	if err != nil {
		return errors.Wrap(err, "write data to json file error")
	}
	fmt.Printf("The file has been generated to %s\n", read_data.OutputFilePath)
	return nil
}
