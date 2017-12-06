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
		ev := even(nums)
		sum += ev
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

func even(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]%nums[j] == 0 {
				return nums[i] / nums[j]
			} else if nums[j]%nums[i] == 0 {
				return nums[j] / nums[i]
			}
		}
	}
	return 0
}
