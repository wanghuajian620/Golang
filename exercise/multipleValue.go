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

	_, c := vals() // _(下划线)是个特殊的变量名，任何赋予它的值都会被丢弃
	fmt.Println(c)
}