// Slices
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

}

//これは最初の要素を含む half-open レンジを選択しますが、最後の要素は省かれます。
//次の式は a の要素の内 1 から 3 を含むスライスを作ります。
// a[1:4]
