package crypto

import (
	"fmt"
	"log"
	"os"

	"github.com/dop251/goja"
)

var (
	vm *goja.Runtime
)

func init() {
	// 读取config
	js_bytes, err := os.ReadFile("assets/bundle.js")
	if err != nil {
		panic(err)
	}
	vm = goja.New()

	// 在Goja VM中执行JavaScript代码
	_, err = vm.RunString(string(js_bytes))
	if err != nil {
		log.Fatalf("Failed to execute JavaScript code: %v", err)
	}
}

func Get_request_url() (string, error) {
	res, err := vm.RunString(`MyLibrary.get_request_url()`)
	return res.String(), err
}
func Parse_result(code string) (string, error) {
	res, err := vm.RunString(fmt.Sprintf(`MyLibrary.get_parsed_result("%v")`, code))
	return res.String(), err
}
