package main

import (
	"fmt"

	. "github.com/sonofachamp/rt-iow/rt"
)

func color(r Ray) Vector {
	unitDirection := UnitVector(r.Direction())
	t := 0.5 * (unitDirection.Y + 1.0)
	return V(1.0, 1.0, 1.0).MultiplyByScalar(1.0 - t).Add(V(0.5, 0.7, 1.0).MultiplyByScalar(t))
}

func main() {
	nx := 1920
	ny := 1080

	fmt.Println("P3")
	fmt.Println(nx, ny)
	fmt.Println(255)

	lowerLeft := V(-2.0, -1.0, -1.0)
	horizontal := V(4.0, 0.0, 0.0)
	vertical := V(0.0, 2.0, 0.0)
	origin := V(0.0, 0.0, 0.0)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			r := NewRay(origin, lowerLeft.Add(horizontal.MultiplyByScalar(u)).Add(vertical.MultiplyByScalar(v)))
			color := color(r)

			ir := int(color.X * 255)
			ig := int(color.Y * 255)
			ib := int(color.Z * 255)

			fmt.Println(ir, ig, ib)
		}
	}
}
