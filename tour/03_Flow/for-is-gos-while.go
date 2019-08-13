//For is Go's "while"
package main

import "fmt"

func main() {
	sum := 1
	// whileはgoではfor.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
