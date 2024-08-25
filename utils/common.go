package utils

import "fmt"



func ToPointer[T string|int|bool|float64](t T) *T {
	return &t
}

func Errorf (format string, a ...interface{}) error {
	f := "error: %s"
	fmt.Printf(f, a...)
	return fmt.Errorf(f, a...)
}