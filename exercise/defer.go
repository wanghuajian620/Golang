/*
 *  Revision History:
 *      Initial: 2018/08/01    Wang Huajian
 */

package main

import (
	"fmt"
	"errors"
)

func deferWithoutReturnName() int {
	var i int

	defer func() {
		i++
		fmt.Println("1st:", i)
	}()

	defer func() {
		i++
		fmt.Println("2nd:", i)
	}()

	return i
}

func deferWithReturnName() (i int) {
	defer func() {
		i++
		fmt.Println("1st:", i)
	}()

	defer func() {
		i++
		fmt.Println("2nd:", i)
	}()

	return i
}

func deferWithAddress() *int  {
	var i int

	defer func() {
		i++
		fmt.Println("1st:", i)
	}()

	defer func() {
		i++
		fmt.Println("2nd:", i)
	}()

	return &i
}

func deferError() error {
	var err error

	defer func() {
		err = errors.New("shouldn't change the return value")
	}()
	return err
}

func main() {
	fmt.Println("deferWithoutReturnName", deferWithoutReturnName())
	fmt.Println("deferWithReturnName", deferWithReturnName())
	fmt.Println("deferWithAddress", *deferWithAddress())
	fmt.Println("deferError", deferError())
}