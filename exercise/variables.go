/*
 *  Revision History:
 *      Initial: 2018/07/23    Wang Huajian
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var 声明一个或多个变量
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

    // go 可以推断出初始变量的类型
	m := 6
	fmt.Println(reflect.TypeOf(m))

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

    // := 是声明和初始化变量的简写
	f:= "short"
	fmt.Println(f)

}
