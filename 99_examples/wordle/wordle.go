package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var words = []string{
	"PLACES", "MARKET", "SUMMER", "OBJECT", "FUTURE", "BUTTON", "NUMBER",
	"PLAYER", "WINTER", "POCKET", "JACKET", "SCHOOL", "WINDOW", "SCREEN",
	"FRIEND",
}

const (
	RightSpot   string = "O"
	WrongSpot          = "X"
	Miss               = "-"
	WordLen            = 6
	MaxAttempts        = 6
)

type Wordle struct {
	attempts    int
	maxAttempts int
	history     []string
	word        string
}

func NewWordle() *Wordle {
	w := &Wordle{}
	w.Reset()
	return w
}

func (w *Wordle) Reset() {
	w.attempts = 0
	w.maxAttempts = MaxAttempts
	w.history = make([]string, 0, WordLen)
	w.word = randomWord()
}

func (w *Wordle) Guess(guess string) (bool, string, error) {
	l := len(guess)
	if l != WordLen {
		return false, "", fmt.Errorf("Word must have exactly %d characters", WordLen)
	}
	guess = strings.ToUpper(guess)
	res := ""
	for i := range WordLen {
		wChar := w.word[i]
		gChar := guess[i]
		switch {
		case wChar == gChar:
			res += RightSpot
		case strings.Contains(w.word, string(gChar)):
			res += WrongSpot
		default:
			res += Miss
		}
	}
	if res == strings.Repeat(RightSpot, WordLen) {
		return true, res, nil
	}
	return false, res, nil
}

func (w *Wordle) AddHistory(guess string, result string) {
	w.history = append(w.history, guess+" "+result)
}

func (w *Wordle) Print() {
	fmt.Println("Attempts:", w.attempts)
	fmt.Println("Max Attempts:", w.maxAttempts)
	for _, v := range w.history {
		fmt.Println(v)
	}
}

func (w *Wordle) Run() {
	for {
		w.Print()
		fmt.Printf("Try a word (length %d):\n", WordLen)
		var guess string
		fmt.Scan(&guess)
		guess = strings.ToUpper(guess)
		won, res, err := w.Guess(guess)
		if err != nil {
			fmt.Println("Invalid word:", err)
			continue
		}
		w.AddHistory(guess, res)
		if won {
			fmt.Println("You Won! Word:", w.word)
			break
		}
		if w.attempts == w.maxAttempts-1 {
			fmt.Println("You Lose! Word:", w.word)
			break
		}
		w.attempts++
	}
}

func randomWord() string {
	i := rand.Intn(len(words))
	return strings.ToUpper(words[i])
}

func main() {
	wordle := NewWordle()
	wordle.Run()
}
