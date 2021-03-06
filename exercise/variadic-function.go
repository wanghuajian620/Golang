/*
 *  Revision History:
 *      Initial: 2018/07/26    Wang Huajian
 */

package main

import (
	"fmt"
)

func sum(nums ...int)  {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
        total += num // total = total + num
	}
	fmt.Println(total)
}

func main()  {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}