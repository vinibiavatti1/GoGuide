package main

import (
	"flag"
	"fmt"
)

var (
	wFlag = flag.Int("w", 60, "width of the automata")
	hFlag = flag.Int("h", 20, "height of the automata")
	rFlag = flag.Int("r", 30, "rule number (0~255) of the automata")
)

type (
	Cell     string
	Automata struct {
		width      int
		height     int
		rule       int
		parsedRule [8]Cell
		cells      [][]Cell
	}
)

const (
	Live Cell = "*"
	Dead Cell = " "
)

func NewAutomata(w, h, rule int) (*Automata, error) {
	a := &Automata{}
	err := a.Resize(w, h)
	if err != nil {
		return nil, fmt.Errorf("resize error: %w", err)
	}
	err = a.SetRule(rule)
	if err != nil {
		return nil, fmt.Errorf("set rule error: %w", err)
	}
	a.Reset()
	return a, nil
}

func (a *Automata) Reset() {
	a.cells = make([][]Cell, a.height)
	for y := range a.height {
		a.cells[y] = make([]Cell, a.width)
		for x := range a.width {
			a.cells[y][x] = Dead
		}
	}
}

func (a *Automata) Resize(w, h int) error {
	if w < 0 || h < 0 {
		return fmt.Errorf("invalid size: w:%d h:%d", w, h)
	}
	a.width = w
	a.height = h
	a.Reset()
	return nil
}

func (a *Automata) Rule() int {
	return a.rule
}

func (a *Automata) SetRule(rule int) error {
	if rule < 0 || rule >= 256 {
		return fmt.Errorf("rule out of range (0~255): %d", rule)
	}
	a.rule = rule
	a.parsedRule = parseRule(rule)
	return nil
}

func (a *Automata) Cell(x, y int) Cell {
	if x < 0 || y < 0 {
		return Dead
	}
	if x >= a.width || y >= a.height {
		return Dead
	}
	return a.cells[y][x]
}

func (a *Automata) SetCell(c Cell, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= a.width || y >= a.height {
		return false
	}
	a.cells[y][x] = c
	return true
}

func (a *Automata) Process() {
	for y := range a.height {
		if y == 0 {
			continue
		}
		for x := range a.width {
			i := 0
			if a.Cell(x+1, y-1) == Live {
				i += 1
			}
			if a.Cell(x-0, y-1) == Live {
				i += 2
			}
			if a.Cell(x-1, y-1) == Live {
				i += 4
			}
			a.cells[y][x] = a.parsedRule[i]
		}
	}
}

func (a *Automata) Render() {
	for y := range a.height {
		for x := range a.width {
			fmt.Print(a.cells[y][x])
		}
		fmt.Println()
	}
}

func parseRule(rule int) [8]Cell {
	r := [8]Cell{}
	for i := range 8 {
		cell := Dead
		if (rule>>i)&1 == 1 {
			cell = Live
		}
		r[i] = cell
	}
	return r
}

func main() {
	flag.Parse()
	a, err := NewAutomata(*wFlag, *hFlag, *rFlag)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	a.SetCell(Live, 30, 0)
	a.Process()
	a.Render()
}
