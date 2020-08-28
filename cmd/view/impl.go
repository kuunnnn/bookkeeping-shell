package view

import (
	"bookkeeping-shell/store"
	"bookkeeping-shell/tool"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os/exec"
	"strings"
)

// 将 web 的图表项目生成到本机目录下
// 将 数据转换成 json 然后生成一个 js 文件
// 打开浏览器
func genViewJS() error {
	// TODO: 这里每次都重新生成了一下是否需要判断下 view 目录存在的话就跳过
	err := tool.GenDistFile(store.FileDirPath);
	if err != nil {
		return errors.Wrap(err, "create web chart file error")
	}
	jsonSlice, err := store.ReadDataToRecordSlice();
	if err != nil {
		return errors.WithStack(err)
	}
	// web chart 里面直接指定 window.BK_DATA 为数据源 所以需要使用 var 或
	output := []string{"var BK_DATA = [\n"}
	for i, _ := range jsonSlice {
		s, _ := json.Marshal(jsonSlice[i])
		output = append(output, string(s), "\n")
		if i != len(jsonSlice)-1 {
			output = append(output, ",\n")
		}
	}
	output = append(output, "];")
	p := fmt.Sprintf("%s/view%s", store.FileDirPath, "/data.js");
	err = ioutil.WriteFile(p, []byte(strings.Join(output, "")), 0644)
	if err != nil {
		return errors.Wrap(err, "write data.js to web chart dir error")
	}
	err = exec.Command(`open`, fmt.Sprintf("%s/view/index.html", store.FileDirPath)).Start()
	if err != nil {
		return errors.Wrap(err, "open browser error")
	}
	return nil
}
