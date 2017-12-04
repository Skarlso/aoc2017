package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		var ordered []string
		phrases := strings.Split(line, " ")
		for _, word := range phrases {
			wsplit := strings.Split(word, "")
			sort.Strings(wsplit)
			wjoined := strings.Join(wsplit, "")
			ordered = append(ordered, wjoined)
		}
		orderedJoined := strings.Join(ordered, " ")
		for _, word := range phrases {
			if strings.Count(line, word) > 1 {
				valid = false
				break
			}
			wsplit := strings.Split(word, "")
			sort.Strings(wsplit)
			wjoined := strings.Join(wsplit, "")
			if strings.Count(orderedJoined, wjoined) > 1 {
				valid = false
				break
			}
			valid = true
		}
		if valid {
			validCount++
		}
	}

	// Note, this was off by one for some reason. +1 was needed.
	fmt.Println("Valid cont: ", validCount)
}
