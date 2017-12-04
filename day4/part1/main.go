package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inFile, _ := os.Open("input.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	validCount := 0
	for scanner.Scan() {
		valid := false
		line := scanner.Text()
		phrases := strings.Split(line, " ")
		for _, word := range phrases {
			if strings.Count(line, word) > 1 {
				valid = false
				break
			}
			valid = true
		}
		if valid {
			validCount++
		}
	}

	fmt.Println("Valid cont: ", validCount)
}
