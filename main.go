//go:generate statik -src=./view/build
// generate 之前还需手动在 html 文件的 head 内添加 <script src="./data.js"></script>
package main

import (
	"bookkeeping-shell/cmd"
)

func main() {
	cmd.Execute()
}
