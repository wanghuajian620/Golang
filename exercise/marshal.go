/*
 *  Revision History:
 *      Initial: 2018/08/03    Wang Huajian
 */

package main

import (
	"fmt"
	"encoding/json"
	"os"
)

func main() {
	// 写一个 ColorGroup 类型
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string // []string 表示是切片类型
	}
	group := ColorGroup{
		ID: 1,
		Name: "Wang Huajian",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	// Marshal 的函数方法 func Marshal(v interface{}) ([]byte, err)
	// Marshal 接收一个参数 类型为 interface{}的值v，返回[]byte类型的值和err类型的值
	b, err :=json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println("b:", b)
}
