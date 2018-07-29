/*
 *  Revision History:
 *      Initial: 2018/07/29    Wang Huajian
 */

package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main()  {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Frad"})
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 30}
	fmt.Println(s.name)

	sp := &s // &的意思是对变量取地址，*是对指针取值。
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}