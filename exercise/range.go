/*
 *  Revision History:
 *      Initial: 2018/07/26    Wang Huajian
 */

package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4} // 切片
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) // 第一个%s和第二个%s都表示字符串，第一个%s对应k值，第二个%s对应v值
	}

	for k:= range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
