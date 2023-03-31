package main

// 1급함수 -> 함수를 다른 변수와 동일하게 다루는 함수
// https://developer.mozilla.org/ko/docs/Glossary/First-class_Function

type calculator func(int, int) int

func main() {
	add := func(i int, j int) int {
		return i + j
	}

	r1 := calc(add, 30, 20)
	println(r1)

	r2 := calc(func(x int, y int) int { return x - y }, 20, 10)
	println(r2)

	println()
	println()
	println()
	println()

}

func calc(f calculator, a int, b int) (result int) {
	result = f(a, b)
	return
}
