package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func abs_int64(x int) int {
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
	steps := strings.Split(strings.TrimSpace(string(content)), ",")
	for _, step := range steps {
		switch step {
		case "n":
			y++
		case "ne":
			x++
		case "se":
			z++
		case "s":
			y--
		case "sw":
			x--
		case "nw":
			z--
		}
	}
	distance := abs_int64(x) + abs_int64(y)
	fmt.Printf("(%d,%d,%d)\n", x, y, z)
	fmt.Println(distance)
}
