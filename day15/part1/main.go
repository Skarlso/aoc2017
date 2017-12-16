package main

import "fmt"

type Generator struct {
	curr   int64
	factor int64
}

func (g *Generator) generateNew() {
	prev := (g.curr * g.factor) % 2147483647
	g.curr = prev
}

func main() {
	genA := Generator{curr: 722, factor: 16807}
	genB := Generator{curr: 354, factor: 48271}
	found := 0
	for i := 0; i < 40000000; i++ {
		genA.generateNew()
		genB.generateNew()
		if (genA.curr & 0xffff) == (genB.curr & 0xffff) {
			found++
		}
	}
	fmt.Println("Found: ", found)
}
