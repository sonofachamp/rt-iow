package main

import (
	"fmt"
)

func main() {
	nx := 200
	ny := 100

	fmt.Println("P3")
	fmt.Println(nx, ny)
	fmt.Println(255)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			rpct := (float64(i) / float64(nx))
			gpct := (float64(j) / float64(ny))
			bpct := .2

			r := int(rpct * 255)
			g := int(gpct * 255)
			b := int(bpct * 255)

			fmt.Println(r, g, b)
		}
	}
}
