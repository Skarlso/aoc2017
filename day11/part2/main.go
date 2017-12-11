package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	x := 0
	y := -1
	z := 1
	steps := strings.Split(string(content), ",")
	max := 0
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
	fmt.Println(x, y, z)
	fmt.Println(max + 1)
}
