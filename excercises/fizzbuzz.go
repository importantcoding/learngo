package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		var f string
		var b string
		if i%3 == 0 {
			f = "Fizz"
		}
		if i%5 == 0 {
			b = "Buzz"
		}

		fmt.Println(i, f+b)
	}
}
