package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile = flag.String("memprofile", "", "write mem profile to file")
)

func main() {
	flag.Parse()

	cpuProfile()
	width := 200
	height := 100

	fmt.Println("P3")
	fmt.Println(width, height)
	fmt.Println(255)

	for row := height - 1; row >= 0; row-- {
		for column := 0; column < width; column++ {
			redPercent := float64(column) / float64(width)
			greenPercent := float64(row) / float64(height)
			bluePercent := .2

			r := int(redPercent * 255)
			g := int(greenPercent * 255)
			b := int(bluePercent * 255)

			fmt.Println(r, g, b)
		}
	}

	memoryProfile()
}

func cpuProfile() {
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
}

func memoryProfile() {
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
