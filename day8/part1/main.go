package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inFile, _ := os.Open("../input.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	registers := make(map[string]int)
	var instructions []string
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " ")
		instructions = append(instructions, line)
		registers[s[0]] = 0
	}
	for _, v := range instructions {
		split := strings.Split(v, " ")
		reg := split[0]
		inst := split[1]
		by, _ := strconv.Atoi(split[2])
		conditionReg := split[4]
		op := split[5]
		check, _ := strconv.Atoi(split[6])
		performOp := false
		switch op {
		case "!=":
			performOp = registers[conditionReg] != check
		case ">":
			performOp = registers[conditionReg] > check
		case "<":
			performOp = registers[conditionReg] < check
		case "==":
			performOp = registers[conditionReg] == check
		case ">=":
			performOp = registers[conditionReg] >= check
		case "<=":
			performOp = registers[conditionReg] <= check
		}
		if performOp {
			if inst == "inc" {
				registers[reg] += by
			} else {
				registers[reg] -= by
			}
		}
	}
	fmt.Println("Max: ", getMax(registers))
}

func getMax(m map[string]int) (max int) {
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return
}
