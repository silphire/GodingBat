// https://codingbat.com/prob/p116624
package main

import "fmt"

func diff21(n int) int {
	if n > 21 {
		return n - 21
	}
	return 21 - n
}

func main() {
	fmt.Printf("f(23) -> 2, actual is %d\n", diff21(23))
	fmt.Printf("f(19) -> 4, actual is %d\n", diff21(23))
}
