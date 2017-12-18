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

var commandMap = map[string]func(*strings.Reader){
	"set": func(lineReader *strings.Reader) {
		var reg string
		var val string
		fmt.Fscanf(lineReader, "set %s %v", &reg, &val)
		set(reg, val)
	},
	"add": func(lineReader *strings.Reader) {
		var reg string
		var val string
		fmt.Fscanf(lineReader, "add %s %v", &reg, &val)
		add(reg, val)
	},
	"mul": func(lineReader *strings.Reader) {
		var reg string
		var val string
		fmt.Fscanf(lineReader, "mul %s %v", &reg, &val)
		mul(reg, val)
	},
	"mod": func(lineReader *strings.Reader) {
		var reg string
		var val string
		fmt.Fscanf(lineReader, "mod %s %v", &reg, &val)
		mod(reg, val)
	},
	"snd": func(lineReader *strings.Reader) {
		var f string
		fmt.Fscanf(lineReader, "snd %s", &f)
		snd(f)
	},
}

var playedFrq = 0

func snd(reg string) {
	playedFrq = register[reg]
	//TODO: Implement playing a sound at a certain frequency
}

func set(reg string, val string) {
	v, err := strconv.Atoi(val)
	if err != nil {
		register[reg] = register[val]
	} else {
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

func jgz(off, val string) (int, bool) {
	var x int
	var y int
	v1, err1 := strconv.Atoi(val)
	if err1 != nil {
		x = register[val]
	} else {
		x = v1
	}
	v2, err2 := strconv.Atoi(off)
	if err2 != nil {
		y = register[off]
	} else {
		y = v2
	}
	if x > 0 {
		return y, true
	}
	return 0, false
}

func main() {
	file := os.Args[1]
	fmt.Println("Input: ", file)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	instructions := strings.Split(string(content), "\n")
	position := 0

	for {
		line := strings.Split(instructions[position], " ")
		lineReader := strings.NewReader(instructions[position])
		if line[0] == "rcv" {
			var reg string
			fmt.Fscanf(lineReader, "rcv %s", &reg)
			if rcv(reg) {
				break
			}
		} else if line[0] == "jgz" {
			var off, val string
			fmt.Fscanf(lineReader, "jgz %v %v", &val, &off)
			if v, ok := jgz(off, val); ok {
				position += v
			} else {
				position++
			}
			continue
		} else {
			commandMap[line[0]](lineReader)
		}
		// fmt.Println(instructions[position])
		position++
	}
	fmt.Println("Frequency: ", playedFrq)
}
