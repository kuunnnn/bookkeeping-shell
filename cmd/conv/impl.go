package conv

import (
	"bookkeeping-shell/store"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

func ToJSON() error {
	jsonSlice, err := store.ReadDataToRecordSlice()
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
	err = ioutil.WriteFile(store.OutputFilePath, []byte(strings.Join(output, "")), 0644)
	if err != nil {
		return errors.Wrap(err, "write data to json file error")
	}
	fmt.Printf("The file has been generated to %s\n", store.OutputFilePath)
	return nil
}
