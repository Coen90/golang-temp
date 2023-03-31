package main

import "fmt"

func main() {
	var a []int
	
	if a == nil {
		fmt.Println(a, " is Nil")
	}
	fmt.Println(len(a), cap(a))

	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s))

	slice := []int{0, 1, 2, 3, 4, 5}
	slice = slice[2:5]  
	fmt.Println(slice) //2,3,4 출력

	slice = append(slice, 6, 7)
	fmt.Println(slice)

	println()
	println()
	println()
	println()
	println()
	println()
	println()

	sliceA := make([] int, 0, 4)

	for i := 1; i <= 50; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}
	fmt.Println(sliceA)

	sliceAA := []int{1, 2, 3}
	sliceBB := []int{4, 5, 6}
	sliceAA = append(sliceAA, sliceBB...)
	fmt.Println(sliceAA)
}