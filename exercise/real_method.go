/*
 *  Revision History:
 *      Initial: 2018/08/24    Wang Huajian
 */

package main

import (
	"fmt"
	"math"
)

type Shape struct {
	width, height  float64
}

type Ball struct {
	radius float64
}

func (r Shape) area() float64 {
	return r.width * r.height
}

func (b Ball) area() float64 {
	return b.radius * b.radius * math.Pi
}

func main() {
	r := Shape{12, 2}
	fmt.Println("Area of r is: ", r.area())
	c := Ball{10}
	fmt.Println("Area of b is: ", c.area())
}
