/*
 *  Revision History:
 *      Initial: 2018/08/28    Wang Huajian
 */

package main

import "fmt"

type Human struct {
	name 	string
	age  	int
	phone 	string
}

type Student struct {
	Human // 匿名字段
	school 		string
	loan		float32
}

type Employee struct {
	Human // 匿名字段
	company		string
	money		float32
}

// Human 对象实现 Sayhi 方法
func (h *Human) Sayhi() {
	fmt.Printf("Hi", h.name, h.phone)
}

// Human 对象实现 Sing 方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("lala,lalala", lyrics)
}

// Human 实现 Guzzle 方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle", beerStein)
}

// Employee 重载 Human 的 Sayhi 方法
func (e *Employee) Sayhi() {
	fmt.Println("Hi, I'm %s, I work at %s, Call me on %s\n", e.name, e.company, e.phone)
}

// Student 实现 BorrowMoney 方法
func (s *Student) BrrowMoney(amount float32) {
	s.loan += amount
}

// Employee 实现 SpendSalary 方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

// Interface Men 被 Human, Student 和 Employee 实现
// 这三个类型都实现了这个接口
type Men interface {
	Sayhi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-xxx"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-xxx"}, "Qinghua", 100}
    sam := Employee{Human{"Sam", 36, "444-222-xxx"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	// 定义 Men 类型的变量
	var i Men

	// i 能存储 Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.Sayhi()
	i.Sing("November rain")

	// i 也能存储 Employee
	i= tom
	fmt.Println("This is tom, an Employee:")
	i.Sayhi()
	i.Sing("Born to be wild")

	// 定义了 slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	// 这三个都是不同类型的元素，但是他们实现了 interface 同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.Sayhi()
	}
}
