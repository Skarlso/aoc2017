package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	// content := "se,se,se,se,n,se,n"
	x := 0
	y := 0
	z := 0
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
	}
	distance := (abs(x) + abs(y) + abs(z)) / 2
	fmt.Println(distance)
}
