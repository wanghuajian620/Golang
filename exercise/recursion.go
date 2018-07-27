/*
 *  Revision History:
 *      Initial: 2018/07/27    Wang Huajian
 */

 // 递归函数
package main

import (
	"fmt"
)

func fact(n int) int {
	if n ==0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}