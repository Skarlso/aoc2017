package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var register = map[string]int{
	"a": 0,
	"b": 0,
	"c": 0,
	"d": 0,
	"e": 0,
	"f": 0,
	"g": 0,
	"h": 0,
	"i": 0,
	"j": 0,
	"k": 0,
	"l": 0,
	"m": 0,
	"n": 0,
	"o": 0,
	"p": 0,
}

var playedFrq = 0

func snd(val int) {
	playedFrq = val
	//TODO: Implement playing a sound at a certain frequency
}

func set(reg string, val string) {
	v, err := strconv.Atoi(val)
	if err != nil {
		register[reg] = register[val]
	} else {
		fmt.Println("Set")
		register[reg] = v
	}
}

func add(reg string, val string) {
	v, err := strconv.Atoi(val)
	if err != nil {
		register[reg] += register[val]
	} else {
		register[reg] += v
	}
}

func mul(reg string, val string) {
	v, err := strconv.Atoi(val)
	if err != nil {
		register[reg] *= register[val]
	} else {
		register[reg] *= v
	}
}

func mod(reg string, val string) {
	v, err := strconv.Atoi(val)
	if err != nil {
		register[reg] %= register[val]
	} else {
		register[reg] %= v
	}
}

func rcv(reg string) bool {
	if register[reg] != 0 {
		return true
	}
	return false
}

func jgz(val string) (ret int) {
	v, err := strconv.Atoi(val)
	if err != nil {
		ret = register[val]
	} else {
		ret = v
	}
	return
}

func main() {
	file := os.Args[1]
	fmt.Println("Input: ", file)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Fscanf
	instructions := strings.Split(string(content), "\n")
	position := 0

	for {
		line := strings.Split(instructions[position], " ")
		lineReader := strings.NewReader(instructions[position])
		switch line[0] {
		case "snd":
			var f int
			fmt.Fscanf(lineReader, "snd %d", &f)
			snd(f)
		case "set":
			var reg string
			var val string
			fmt.Fscanf(lineReader, "set %s %v", &reg, &val)
			set(reg, val)
			os.Exit(1)
		case "add":
			var reg string
			var val string
			fmt.Fscanf(lineReader, "add %s %v", &reg, &val)
			add(reg, val)
		case "mul":
			var reg string
			var val string
			fmt.Fscanf(lineReader, "mul %s %v", &reg, &val)
			mul(reg, val)
		case "mod":
			var reg string
			var val string
			fmt.Fscanf(lineReader, "mod %s %v", &reg, &val)
			mod(reg, val)
		case "rcv":
			var reg string
			fmt.Fscanf(lineReader, "rcv %s", &reg)
			if rcv(reg) {
				break
			}
		case "jgz":
			var val string
			fmt.Fscanf(lineReader, "jgz %v", &val)
			position += jgz(val)
		default:
			break
		}
		fmt.Println(instructions[position])
	}
	fmt.Println("Frequency: ", playedFrq)
}
