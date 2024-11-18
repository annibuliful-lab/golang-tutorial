package pkg

import (
	"backend/basic/number"
	"errors"
)

func Add() int {
	return number.A + 2
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
