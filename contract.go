package main


// TODO: create struct to this contract
type hangmanContract interface {
	usedLetter(word, opt string) bool
	operations(ch chan int, opt string)
	userLooser() bool
}
