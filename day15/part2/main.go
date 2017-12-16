package main

import "fmt"

type Generator struct {
	curr   int64
	factor int64
	mult   int64
}

func (g *Generator) generateNew() {
	for {
		g.curr = (g.curr * g.factor) % 2147483647
		if g.curr%g.mult == 0 {
			break
		}
	}
}

func main() {
	genA := Generator{curr: 722, factor: 16807, mult: 4}
	genB := Generator{curr: 354, factor: 48271, mult: 8}
	found := 0
	for i := 0; i < 5000000; i++ {
		genA.generateNew()
		genB.generateNew()
		if (genA.curr & 0xffff) == (genB.curr & 0xffff) {
			found++
		}
	}
	fmt.Println("Found: ", found)
}
