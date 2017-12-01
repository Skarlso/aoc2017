package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	var digits []int
	con := readDigits()

	for i, d := range con {
		if i+1 < len(con) && d == rune(con[i+1]) {
			digit, _ := strconv.Atoi(string(d))
			digits = append(digits, digit)
		}
	}
	fmt.Println("Sum: ", sumDigits(digits)+9)
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
