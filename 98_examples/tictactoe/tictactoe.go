package main

import (
	"fmt"
	"strings"
)

const (
	Len int    = 3
	Div string = " | "
)

type Player string

const (
	Player1 Player = "X"
	Player2 Player = "O"
)

type GameResult int

const (
	None GameResult = iota
	Won
	Tie
)

var positionMap = map[string][2]int{
	"1": [...]int{0, 0},
	"2": [...]int{0, 1},
	"3": [...]int{0, 2},
	"4": [...]int{1, 0},
	"5": [...]int{1, 1},
	"6": [...]int{1, 2},
	"7": [...]int{2, 0},
	"8": [...]int{2, 1},
	"9": [...]int{2, 2},
}

var winningCombinations = [][Len]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
	{"1", "4", "7"},
	{"2", "5", "8"},
	{"3", "6", "9"},
	{"1", "5", "9"},
	{"3", "5", "7"},
}

type Game struct {
	board         [Len][Len]string
	currentPlayer Player
}

func NewGame() *Game {
	game := &Game{
		currentPlayer: Player1,
	}
	game.Reset()
	return game
}

func (g *Game) Reset() {
	g.board = [Len][Len]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
}

func (g *Game) Play(pos string) error {
	coords, ok := positionMap[pos]
	if !ok {
		return fmt.Errorf("Invalid Position: %q", pos)
	}
	x, y := coords[0], coords[1]
	boardpos := g.board[x][y]
	if boardpos == string(Player1) || boardpos == string(Player2) {
		return fmt.Errorf("Position is not empty")
	}
	g.board[x][y] = string(g.currentPlayer)
	return nil
}

func (g *Game) SwapCurrentPlayer() {
	if g.currentPlayer == Player1 {
		g.currentPlayer = Player2
		return
	}
	g.currentPlayer = Player1
}

func (g *Game) ValidateWinner() bool {
	for _, v := range winningCombinations {
		if g.IsSamePlayerInPositions(v[0], v[1], v[2]) {
			return true
		}
	}
	return false
}

func (g *Game) ValidateTie() bool {
	for i := range g.board {
		for j := range g.board {
			pos := g.board[i][j]
			if pos != string(Player1) && pos != string(Player2) {
				return false
			}
		}
	}
	return true
}

func (g *Game) Validate() GameResult {
	switch {
	case g.ValidateWinner():
		return Won
	case g.ValidateTie():
		return Tie
	}
	return None
}

func (g *Game) IsSamePlayerInPositions(x, y, z string) bool {
	x1, y1 := positionMap[x][0], positionMap[x][1]
	x2, y2 := positionMap[y][0], positionMap[y][1]
	x3, y3 := positionMap[z][0], positionMap[z][1]
	return g.board[x1][y1] == g.board[x2][y2] && g.board[x2][y2] == g.board[x3][y3]
}

func (g *Game) PrintBoard() {
	for i := range g.board {
		line := strings.Join(g.board[i][:], Div)
		fmt.Println(line)
	}
}

func (g *Game) Run() {
	running := true
	for running {
		fmt.Println("Current Player:", g.currentPlayer)
		g.PrintBoard()
		for {
			pos := Scan()
			err := g.Play(pos)
			if err != nil {
				fmt.Println(err)
			} else {
				break
			}
		}
		result := g.Validate()
		switch result {
		case None:
			g.SwapCurrentPlayer()
		case Tie:
			fmt.Println("Tie!!!")
			running = false
		case Won:
			fmt.Printf("Player %s Won!\n", g.currentPlayer)
			running = false
		}
	}
	fmt.Println("Game finished")
	g.PrintBoard()
}

func Scan() string {
	for {
		var pos string
		fmt.Print("Position: ")
		_, err := fmt.Scanln(&pos)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return pos
	}
}

func main() {
	game := NewGame()
	game.Run()
}
