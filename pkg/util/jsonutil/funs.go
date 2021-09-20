package jsonutil

import (
	"encoding/json"
	"fmt"
)

//json方式打印结构体
func PrintJson(obj interface{}) {
	tmp, _ := json.MarshalIndent(obj, "", "     ")
	fmt.Println(string(tmp))
}

