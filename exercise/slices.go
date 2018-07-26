/*
 *  Revision History:
 *      Initial: 2018/07/25    Wang Huajian
 */

package main

import "fmt"

func main() {
    // 使用内置的make方法创建一个非零长度的切片
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd", s)

	c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // 可以得到切片中s[2],s[3],s[4]的值
    l := s[2:5]
    fmt.Println("sl1:", l)
    l = s[:5]
    fmt.Println("sl2:", l)
    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // 切片可以组成多维数组， 但是不同于数组的是内部切片的长度可以不同
    // 比如 [[0][1 2][2 3 4]]
    twoD := make([][]int, 3)
    fmt.Println("o:", twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
