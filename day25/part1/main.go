package main

import "fmt"

type State struct {
	name      string
	value     int
	behaviour func(int) (int, int, *State)
}

type Tape struct {
	cursor       int
	tape         []int
	currentState *State
}

func (t *Tape) extend() {
	t.tape = append(t.tape, 0)
}

func (t *Tape) prepend() {
	tp := make([]int, 0, 0)
	tp = append(tp, 0)
	t.tape = append(tp, t.tape...)
	t.cursor++
}

func countOnes(tape []int) (sum int) {
	for _, v := range tape {
		sum += v
	}
	return
}

func main() {
	steps := 12523873
	a := State{name: "A", value: 0}
	b := State{name: "B", value: 0}
	c := State{name: "C", value: 0}
	d := State{name: "D", value: 0}
	e := State{name: "E", value: 0}
	f := State{name: "F", value: 0}
	a.behaviour = func(readValue int) (writeValue int, cursor int, state *State) {
		if readValue == 0 {
			return 1, 1, &b
		} else if readValue == 1 {
			return 1, -1, &e
		}
		return 0, 0, nil
	}
	b.behaviour = func(readValue int) (writeValue int, cursor int, state *State) {
		if readValue == 0 {
			return 1, 1, &c
		} else if readValue == 1 {
			return 1, 1, &f
		}
		return 0, 0, nil
	}
	c.behaviour = func(readValue int) (writeValue int, cursor int, state *State) {
		if readValue == 0 {
			return 1, -1, &d
		} else if readValue == 1 {
			return 0, 1, &b
		}
		return 0, 0, nil
	}
	d.behaviour = func(readValue int) (writeValue int, cursor int, state *State) {
		if readValue == 0 {
			return 1, 1, &e
		} else if readValue == 1 {
			return 0, -1, &c
		}
		return 0, 0, nil
	}
	e.behaviour = func(readValue int) (writeValue int, cursor int, state *State) {
		if readValue == 0 {
			return 1, -1, &a
		} else if readValue == 1 {
			return 0, 1, &d
		}
		return 0, 0, nil
	}
	f.behaviour = func(readValue int) (writeValue int, cursor int, state *State) {
		if readValue == 0 {
			return 1, 1, &a
		} else if readValue == 1 {
			return 1, 1, &c
		}
		return 0, 0, nil
	}
	t := make([]int, 0, 0)
	t = append(t, 0)
	tape := Tape{cursor: 0, tape: t, currentState: &a}

	for i := 0; i < steps; i++ {
		writeValue, cursor, state := tape.currentState.behaviour(tape.tape[tape.cursor])
		tape.tape[tape.cursor] = writeValue
		tape.cursor += cursor
		if tape.cursor >= len(tape.tape) {
			tape.extend()
		} else if tape.cursor <= 0 {
			tape.prepend()
		}
		tape.currentState = state
	}

	fmt.Println("Sum of 1s: ", countOnes(tape.tape))
}
