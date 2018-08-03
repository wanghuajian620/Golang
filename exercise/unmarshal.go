/*
 *  Revision History:
 *      Initial: 2018/08/03    Wang Huajian
 */

package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	// 先写一个[]byte类型，并赋值给jsonBlob
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)

	// 声明一个新的类型
	type Animal struct {
		Name string
		Order string
	}

	var animals []Animal // 再声明一个animals 为 []Animal类型
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
