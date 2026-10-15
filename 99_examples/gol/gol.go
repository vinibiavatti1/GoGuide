package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
	"time"
)

type (
	Cell string
	Rule struct {
		Born    []int
		Survive []int
	}
	GameOfLife struct {
		Width  int
		Height int
		Rule   *Rule
		screen [][]Cell
	}
	update struct {
		x    int
		y    int
		cell Cell
	}
)

const (
	Live Cell = "*"
	Dead Cell = " "
)

var (
	Gol, _          = ParseRule("3/23")
	HighLife, _     = ParseRule("23/36")
	Assimilation, _ = ParseRule("345/4567")
	TwoXTwo, _      = ParseRule("36/125")
	DayAndNight, _  = ParseRule("3678/34578")
	Amoeba, _       = ParseRule("357/1358")
	Move, _         = ParseRule("368/245")
	PseudoLife, _   = ParseRule("357/238")
	Diamoeba, _     = ParseRule("35678/5678")
	Rule34, _       = ParseRule("34/34")
	LongLife, _     = ParseRule("345/5")
	Stains, _       = ParseRule("3678/235678")
	Seeds, _        = ParseRule("2/")
	Maze, _         = ParseRule("3/12345")
	Coagulations, _ = ParseRule("378/235678")
	WalledCities, _ = ParseRule("45678/2345")
	Gnarl, _        = ParseRule("1/1")
	Replicator, _   = ParseRule("1357/1357")
	Mystery, _      = ParseRule("3458/05678")
)

func NewGameOfLife(w, h int, rule *Rule) *GameOfLife {
	gol := &GameOfLife{
		Width:  w,
		Height: h,
		Rule:   rule,
	}
	gol.Reset()
	return gol
}

func (g *GameOfLife) Reset() {
	g.screen = make([][]Cell, g.Height)
	for y := range g.Height {
		g.screen[y] = make([]Cell, g.Width)
		for x := range g.Width {
			g.screen[y][x] = Dead
		}
	}
}

func (g *GameOfLife) Get(x, y int) Cell {
	if x < 0 || y < 0 {
		return Dead
	}
	if x >= g.Width || y >= g.Height {
		return Dead
	}
	return g.screen[y][x]
}

func (g *GameOfLife) Set(c Cell, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= g.Width || y >= g.Height {
		return false
	}
	g.screen[y][x] = c
	return true
}

func (g *GameOfLife) Render() {
	for y := range g.Height {
		for x := range g.Width {
			c := g.Get(x, y)
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}

func (g *GameOfLife) Next() {
	updates := []*update{}
	for y := range g.Height {
		for x := range g.Width {
			cell := g.Get(x, y)
			lives := g.countLives(x, y)
			if cell == Live && !slices.Contains(g.Rule.Survive, lives) {
				updates = append(updates, &update{x: x, y: y, cell: Dead})
			} else if slices.Contains(g.Rule.Born, lives) {
				updates = append(updates, &update{x: x, y: y, cell: Live})
			}
		}
	}
	for _, v := range updates {
		g.Set(v.cell, v.x, v.y)
	}
}

func (g *GameOfLife) countLives(x, y int) int {
	cells := [8]Cell{
		g.Get(x-1, y-1),
		g.Get(x-0, y-1),
		g.Get(x+1, y-1),
		g.Get(x-1, y-0),
		g.Get(x+1, y-0),
		g.Get(x-1, y+1),
		g.Get(x-0, y+1),
		g.Get(x+1, y+1),
	}
	count := 0
	for _, c := range cells {
		if c == Live {
			count++
		}
	}
	return count
}

func ParseRule(rule string) (*Rule, error) {
	before, after, found := strings.Cut(rule, "/")
	if !found {
		return nil, fmt.Errorf("missing %q separator", "/")
	}
	born := []int{}
	survive := []int{}
	for _, v := range before {
		if i, err := strconv.Atoi(string(v)); err == nil {
			born = append(born, i)
		}
	}
	for _, v := range after {
		if i, err := strconv.Atoi(string(v)); err == nil {
			survive = append(survive, i)
		}
	}
	return &Rule{
		Born:    born,
		Survive: survive,
	}, nil
}

func run(cmd string) {
	c := exec.Command("cmd", "/c", cmd)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}

func main() {
	gol := NewGameOfLife(60, 20, Gol)
	gol.Set(Live, 30, 8)
	gol.Set(Live, 30, 9)
	gol.Set(Live, 30, 10)
	gol.Set(Live, 29, 9)
	gol.Set(Live, 31, 10)
	run("@echo off && cls")
	for {
		gol.Next()
		gol.Render()
		time.Sleep(time.Millisecond * 100)
		run("cls")
	}
}
