/*
 *  Revision History:
 *      Initial: 2018/07/27    Wang Huajian
 */

// 告诉我们当前文件属于哪个包，而包名main告诉我们它是一个可独立运行的包，它在编译后会产生可执行文件
// 除了main包之外，其他的包都会生成*.a文件（也就是包文件）并放置在$GOPATH/pkg/$GOOS_$GOARCH中（以Mac为例就是$GOPATH/pkg/darwin_amd64）
// 每一个可独立运行的Go程序，必定包含一个package main, 在这个main包中必定包含一个入口函数main，而这个函数既没有参数，也没有返回值。
package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0 // := 这个符号直接取代了var和type，但是它有一个限制就是必须在函数内部使用；在函数外部使用则会无法编译通过，所以一般使用var方式来定义全局变量
	return func() int {
		i++
		return i
	}
}

func main()  {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
