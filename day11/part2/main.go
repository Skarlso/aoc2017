package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	content, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	x := 0
	y := 0
	z := 0
	max := 0
	steps := strings.Split(strings.TrimSpace(string(content)), ",")
	for _, step := range steps {
		switch step {
		case "n":
			x--
			y++
		case "ne":
			y++
			z--
		case "se":
			x++
			z--
		case "s":
			x++
			y--
		case "sw":
			y--
			z++
		case "nw":
			x--
			z++
		}
		max = int(math.Max(float64(max), math.Abs(float64(x))))
		max = int(math.Max(float64(max), math.Abs(float64(y))))
		max = int(math.Max(float64(max), math.Abs(float64(z))))
	}
	distance := (abs(x) + abs(y) + abs(z)) / 2
	fmt.Println(distance)
	fmt.Println(max)
}
