package main


// TODO: create struct to this contract
type HangmanContract interface {
	sortWord()
	isLooser() bool
	showBoard()
	clearTerminal()
	showWord(isStart bool)
	usedLetter(word, opt string) bool
	operations(ch chan int, opt string)
}
