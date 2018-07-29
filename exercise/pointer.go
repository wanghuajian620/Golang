/*
 *  Revision History:
 *      Initial: 2018/07/27    Wang Huajian
 */

package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) { // 这里的*int是说明 iptr 是指针类型
	*iptr = 0
}

func main()  {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i) // 结果为1，是因为i的值1没有被zeroval这个函数改变，只有传地址才可以改变。

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}