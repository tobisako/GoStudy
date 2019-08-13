// Constants
package main

import "fmt"

const Pi = 3.14

func main() {
	// なお、定数は := を使って宣言できません。
	const World = "世界"

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
