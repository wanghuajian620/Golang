/*
 *  Revision History:
 *      Initial: 2018/07/24    Wang Huajian
 */

package main

import (
	"fmt"
)

func main() {
	// 元素的类型和长度都是数组类型的一部分。默认情况下，数组是零值的，ints表示0。
	var a[5] int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a) // why
	fmt.Println("get:", a[4])
	fmt.Println("length:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 数组类型是唯一的， 但是我们可以组合类型来构建多维数组结构 【【0 1 2】【1 2 3】】
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
