// Exercise: Slices
package main

import "golang.org/x/tour/pic"

//import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for y := range image {
		image[y] = make([]uint8, dx)
	}
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			image[y][x] = uint8((x + y) / 2)
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}

// Answer
// https://qiita.com/kroton/items/8622d5e9e38ff822070c
