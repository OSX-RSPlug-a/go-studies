package main

import "fmt"

func fibonacci() func() int {
	val1 := 0
	val2 := 1

	return func() int {
		defer func() { val1, val2 = val2, val1+val2 }()
		return val1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
