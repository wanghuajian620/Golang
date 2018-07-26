/*
 *  Revision History:
 *      Initial: 2018/07/23    Wang Huajian
 */

package main

import (
	"fmt"
	"math"
)

// go 支持字符、字符串、布尔值和数值的常量
// const 声明一个常量的值
const s string = "constants"

func main() {
    fmt.Println(s)

    const n = 500000000

    // 常量表达式执行任意精度的算术运算。
    const d = 3e20 / n
    fmt.Println(d)

    // 数字常量在给定类型前没有类型，我们可以通过显式转换
    fmt.Println(int64(d))
    fmt.Println(math.Sin(n))
}
