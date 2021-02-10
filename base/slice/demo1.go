package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	sliceAdd(a)
	fmt.Println(a)
}

func sliceAdd(s []int)  {
	s = append(s, 4)
}
