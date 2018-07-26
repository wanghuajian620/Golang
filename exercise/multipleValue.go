/*
 *  Revision History:
 *      Initial: 2018/07/26    Wang Huajian
 */

package main

import (
	"fmt"
)

// 多个返回值的函数
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}