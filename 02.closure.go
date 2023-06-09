package main

func main() {

	next := nextValue()
	println(next())
	println(next())
	println(next())
	println(next())

	anotherNext := nextValue()

	println(anotherNext())
	println(anotherNext())

	println(next())
	println(next())
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
