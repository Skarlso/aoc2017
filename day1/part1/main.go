package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	var digits []int
	con := readDigits()
	l := len(con)
	for i, d := range con {
		if d == rune(con[(i+1)%l]) {
			digit, _ := strconv.Atoi(string(d))
			digits = append(digits, digit)
		}
	}
	fmt.Println("Sum: ", sumDigits(digits))
}

func readDigits() string {
	con, _ := ioutil.ReadFile("input.txt")
	return string(con)
}

func sumDigits(digits []int) (sum int) {
	for _, d := range digits {
		sum += d
	}
	return
}
