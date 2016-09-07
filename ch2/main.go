package main

import (
	"fmt"

	. "github.com/sonofachamp/rt-iow/rt"
)

func main() {
	nx := 200
	ny := 100

	fmt.Println("P3")
	fmt.Println(nx, ny)
	fmt.Println(255)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			color := V(float64(i)/float64(nx), float64(j)/float64(ny), 0.2)

			r := int(color.X * 255)
			g := int(color.Y * 255)
			b := int(color.Z * 255)

			fmt.Println(r, g, b)
		}
	}
}
