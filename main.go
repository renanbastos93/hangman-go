package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	words      []string
	board      []string
	posBoard   int
	chance     int
	selectWord string
	userWrite  string
	wrongWord  string
	isLooser   bool
	out        []string
)

func createBoard() {
	re := regexp.MustCompile(`[\d]`)
	steps := []string{"O", "|", "/", "\\", "/", "\\"}
	base := `
 +----+
 |    |
 1    |
324   |
5 6   |
      |
==========`

	board = make([]string, len(steps)+1)
	board[0] = re.ReplaceAllString(base, " ")
	for i, s := range steps {
		base = strings.Replace(base, strconv.Itoa(i+1), s, 1)
		currentBase := re.ReplaceAllString(base, " ")
		board[i+1] = currentBase
	}
}

func sortWord() {
	words = []string{"bird", "happy", "soccer", "computer"}
	rand.Seed(time.Now().UnixNano())
	selectWord = words[rand.Intn(len(words)-1)+1]
}

func showBoard(idx int) {
	if idx < len(board) {
		fmt.Println(board[idx])
	} else {
		isLooser = true
	}
}

func usedLetter(word, opt string) bool {
	for _, v := range word {
		if v == []rune(opt)[0] {
			return true
		}
	}
	return false
}

func operations(ch chan int, opt string) {
	countWord := strings.Count(selectWord, opt)
	if strings.Contains(selectWord, opt) && (strings.Count(wrongWord, opt) >= countWord ||
		strings.Count(userWrite, opt) >= countWord) || usedLetter(wrongWord, opt) || usedLetter(userWrite, opt) {
		chance--
		posBoard++
		wrongWord = wrongWord + opt
		showBoard(posBoard)
	} else {
		for i, v := range selectWord {
			if v == []rune(opt)[0] {
				out[i] = opt
			}
		}
	}
	userWrite = userWrite + opt
	ch <- 1
}

func userLooser() bool {
	return (chance == 0 || isLooser == true) && strings.Join(out, "") != selectWord
}

func run() {
	var input string
	ch := make(chan int, 1)
	chance = 6
	posBoard = 0
	isLooser = false
	out = make([]string, len(selectWord))
	for i := range selectWord {
		out[i] = "_"
	}
	fmt.Println(selectWord, "\n", board[posBoard])
	for {
		fmt.Println("Word: ", out)
		if userLooser() {
			fmt.Println("You are a loser...")
			return
		}
		if strings.Join(out, "") == selectWord {
			fmt.Println("You are a winner...")
			return
		}
		fmt.Scanln(&input)
		if len(input) == 1 && unicode.IsLetter([]rune(input)[0]) {
			operations(ch, input)
			<-ch
		} else {
			if input == "exit" {
				return
			}
			fmt.Println("parameter not permission")
		}
		input = ""
	}
}

func main() {
	sortWord()
	createBoard()
	run()
}
