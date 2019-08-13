// Stacking defers
package main

import "fmt"

// returnするとき、
// defer へ渡した関数は LIFO(last-in-first-out) の順番で実行されます。

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
