/*
 *  Revision History:
 *      Initial: 2018/07/31    Wang Huajian
 */

package main

import (
	"errors"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil  // A nil value in the error position indicates that there was no error.
}

type argError struct {
	arg int
	prob string
}
