/*
 *  Revision History:
 *      Initial: 2018/07/30    Wang Huajian
 */

package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	weidth, height float64
}

type Circle struct {
	radius float64
}

// method的语法：func (r ReceiverType) funcName(parameters) (results)
// 是方法而不是函数
func (r Rectangle) area() float64 {
	return r.weidth*r.height
}

func (c Circle) area() float64 {
    return c.radius*c.radius*math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}