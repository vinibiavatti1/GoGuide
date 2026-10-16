package main

import (
	"flag"
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

var (
	difficultyFlag  = flag.Int("d", 3, "The difficulty (1-Very Easy, 2-Easy, 3-Medium, 4-Hard, 5-Very Hard)")
	customFlag      = flag.Bool("c", false, "Set as custom game. Must set the flags: (-n -l -a -r)")
	digitsFlag      = flag.Int("cn", 4, "(Custom game) The number of allowed digits")
	lengthFlag      = flag.Int("cl", 4, "(Custom game) The length of the password")
	attemptsFlag    = flag.Int("ca", 4, "(Custom game) The max number of attempts")
	repetitionsFlag = flag.Bool("cr", false, "(Custom game) Allow digit repetitions")
)

const (
	FirstDigitCodePoint = 65 // 'A'
	LastDigitCodePoint  = 90 // 'Z'
	MaxDigits           = LastDigitCodePoint - FirstDigitCodePoint
)

type Difficulty int

const (
	Custom = iota
	VeryEasy
	Easy
	Medium
	Hard
	VeryHard
)

var difficultyNameMap = map[Difficulty]string{
	VeryEasy: "Very Easy",
	Easy:     "Easy",
	Medium:   "Medium",
	Hard:     "Hard",
	VeryHard: "Very Hard",
	Custom:   "Custom",
}

func (d *Difficulty) String() string {
	name, ok := difficultyNameMap[*d]
	if !ok {
		return "Unknown"
	}
	return name
}

type guess struct {
	input string
	has   int
	hit   int
	total int
}

func (g *guess) String() string {
	guess := ""
	for _, c := range g.input {
		guess += string(c)
	}
	return guess
}

type Mastermind struct {
	maxAttempts      int
	attempts         int
	allowRepetitions bool
	passwordLen      int
	digits           string
	password         string
	guesses          []*guess
	difficulty       Difficulty
}

func New(nDigits, passwordLen, maxAttempts int, allowRepetitions bool) (*Mastermind, error) {
	if passwordLen <= 0 {
		return nil, fmt.Errorf("password length must be > 0")
	}
	if nDigits <= 0 || nDigits > MaxDigits {
		return nil, fmt.Errorf("nDigits must be 1~%d", MaxDigits)
	}
	if !allowRepetitions && passwordLen > nDigits {
		return nil, fmt.Errorf("password length must be between 1~%d", nDigits)
	}
	m := &Mastermind{
		allowRepetitions: allowRepetitions,
		passwordLen:      passwordLen,
		maxAttempts:      maxAttempts,
		digits:           generateDigits(nDigits),
		difficulty:       Custom,
	}
	m.Reset()
	return m, nil
}

func NewFromDifficulty(d Difficulty) (*Mastermind, error) {
	var m *Mastermind
	var err error
	switch d {
	case VeryEasy:
		m, err = New(5, 3, 10, false)
	case Easy:
		m, err = New(6, 3, 10, false)
	case Medium:
		m, err = New(6, 4, 10, false)
	case Hard:
		m, err = New(8, 5, 12, false)
	case VeryHard:
		m, err = New(8, 5, 12, true)
	default:
		return nil, fmt.Errorf("invalid difficulty level: %d", d)
	}
	m.difficulty = d
	return m, err
}

func (m *Mastermind) Reset() {
	m.attempts = 0
	m.password = randomPassword(m.digits, m.passwordLen, m.allowRepetitions)
	m.guesses = make([]*guess, 0, m.maxAttempts)
}

func (m *Mastermind) Print() {
	fmt.Printf("%-3s %-11s %s\n", "#", "Password", "Hit/Has/Tot")
	for i := range m.maxAttempts {
		attempt := i + 1
		if i < len(m.guesses) {
			g := m.guesses[i]
			fmt.Printf("%02d  %-11s %d/%d/%d\n", attempt, g.String(), g.hit, g.has, g.total)
			continue
		}
		fmt.Printf("%02d  %-11s %s\n", attempt, strings.Repeat("-", m.passwordLen), "-/-/-")
	}
	fmt.Println()
}

func (m *Mastermind) Guess(input string) (bool, error) {
	input = strings.ToUpper(input)
	if len(input) != m.passwordLen {
		return false, fmt.Errorf("input must have exactly %d digits", m.passwordLen)
	}
	guess := &guess{}
	for i, v := range input {
		digit := string(v)
		if !strings.Contains(m.digits, string(digit)) {
			return false, fmt.Errorf("input has invalid characters: %s", digit)
		}
		if m.password[i] == byte(v) {
			guess.hit++
		} else if strings.Contains(m.password, digit) {
			guess.has++
		}
		guess.input += digit
	}
	guess.total = guess.hit + guess.has
	m.guesses = append(m.guesses, guess)
	if guess.hit == m.passwordLen {
		return true, nil
	}
	return false, nil
}

func (m *Mastermind) DigitsString() string {
	res := ""
	for _, v := range m.digits {
		res += string(v)
	}
	return res
}

func (m *Mastermind) PasswordString() string {
	res := ""
	for _, v := range m.password {
		res += string(v)
	}
	return res
}

func (m *Mastermind) Run() {
	fmt.Printf("Welcome to Mastermind!\n")
	fmt.Printf("Difficulty: %s\n", m.difficulty.String())
	fmt.Printf("Digits: %v\n", m.DigitsString())
	fmt.Printf("Password Len: %d\n", m.passwordLen)
	fmt.Printf("Allow Repetition: %v\n", m.allowRepetitions)
	fmt.Printf("Max Attempts: %d\n", m.maxAttempts)
	m.Print()
	for {
		fmt.Printf("Try a Password (Digits: %v / Len: %d): ", m.DigitsString(), m.passwordLen)
		var password string
		fmt.Scan(&password)
		won, err := m.Guess(password)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		m.Print()
		if won {
			fmt.Printf("You Won! (Password: %s)\n", m.PasswordString())
			return
		}
		m.attempts++
		if m.attempts == m.maxAttempts {
			fmt.Printf("You Lose! (Password: %s)\n", m.PasswordString())
			return
		}
	}
}

func generateDigits(nDigits int) string {
	res := ""
	for i := range nDigits {
		res += string(rune(i + FirstDigitCodePoint))
	}
	return res
}

func randomPassword(digits string, passwordLen int, allowRepetitions bool) string {
	availableDigits := strings.Split(digits, "")
	password := ""
	for range passwordLen {
		rnd := rand.Intn(len(availableDigits))
		digit := availableDigits[rnd]
		password += string(digit)
		if !allowRepetitions {
			availableDigits = slices.Delete(availableDigits, rnd, rnd+1)
		}
	}
	return password
}

func main() {
	flag.Parse()
	var m *Mastermind
	var err error
	if !*customFlag {
		m, err = NewFromDifficulty(Difficulty(*difficultyFlag))
	} else {
		m, err = New(*digitsFlag, *lengthFlag, *attemptsFlag, *repetitionsFlag)
	}
	if err != nil {
		panic(err)
	}
	m.Run()
}
