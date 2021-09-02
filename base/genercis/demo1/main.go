package main

import "fmt"

type comparable interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64
}

func main() {
	f1(1)
}

func f1[T any](v1 T)  {
	f2(v1)
}

func f2[T any](v1 T)  {
	fmt.Println(v1)
}

type A[T any] struct {
	a int
	b T
}