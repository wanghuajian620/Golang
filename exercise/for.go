/*
 *  Revision History:
 *      Initial: 2018/07/23    Wang Huajian
 */

package main

import "fmt"

func main() {

	// 大多数基本的类型，只有一个条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

    // 经典的情况：赋值 条件 循环
	for j := 7; j <= 9 ; j ++ {
		fmt.Println(j)
	}

	for  {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++  {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}