/*
 *  Revision History:
 *      Initial: 2018/07/23    Wang Huajian
 */

package main

import (
	"fmt"
)

func main() {
	var s = 7%2
	fmt.Println(s)

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 可以被4整除")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
