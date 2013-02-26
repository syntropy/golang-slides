package main

import "fmt"

func doAndPrint(x, y string, f func(x, y string) (string, string)) {
	x, y = f(x, y)
	fmt.Println(x, y)
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	doAndPrint("x", "y", swap)
}
