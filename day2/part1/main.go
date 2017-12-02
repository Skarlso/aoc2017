package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inFile, _ := os.Open("input.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	var reg = regexp.MustCompile(`\d+`)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := reg.FindAllString(line, -1)
		nums := convert(numbers)
		mi := min(nums)
		ma := max(nums)
		sum += ma - mi
	}
	fmt.Println("Sum: ", sum)
}

func convert(nums []string) (ret []int) {
	for _, n := range nums {
		i, _ := strconv.Atoi(n)
		ret = append(ret, i)
	}
	return
}

func min(nums []int) int {
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}

	return m
}

func max(nums []int) int {
	m := 0
	for _, n := range nums {
		if n > m {
			m = n
		}
	}

	return m
}
