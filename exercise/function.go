/*
 *  Revision History:
 *      Initial: 2018/07/26    Wang Huajian
 */

package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

// 第一个int给 a,b,c 定义类型， 第二个int是给返回值定义类型
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1,2, 3)
	fmt.Println("1+2+3=", res)
}