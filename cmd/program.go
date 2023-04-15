package main

import "fmt"

type Program struct {
	lines []Line
	rows  int
}

func NewProgram() *Program {
	return &Program{rows: 1, lines: []Line{{
		source: []Token{},
		length: 0,
	}}}
}

func (p Program) Debug() {
	for _, l := range p.lines {
		for _, t := range l.source {
			t.Print()
		}
		fmt.Printf("\n")
	}
}
