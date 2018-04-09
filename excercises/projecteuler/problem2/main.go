package main

import "fmt"

func main() {
	var pa = 1
	var pr = 1
	var sum int
	var total int
	for pr <= 4000000 {
		sum = pa + pr
		fmt.Println(sum)
		pa = pr
		pr = sum
		if sum%2 == 0 {
			total += sum
		}
	}

	fmt.Println(total)
}

//https://projecteuler.net/problem=2

/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/
