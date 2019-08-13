// Slice length and capacity
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice
	s = s[:0]
	printSlice(s)

	// Extend
	s = s[:4]
	printSlice(s)

	// Drop
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// 十分な容量を持って提供されているスライスを再スライスすることによって、
// スライスの長さを伸ばすことができます。
// その容量を超えて伸ばしたときに何が起こるかを見るため、
// プログラム例でのスライスのいずれかの操作を変更してみてください。
