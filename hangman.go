package main

type HangmanGame struct {
	Words      []string
	Board      []string
	PosBoard   int
	Chance     int
	SelectWord string
	UserWrite  string
	WrongWord  string
	IsLooser   bool
	Out        []string
}

const ClearTerminal = "\033[2J"

var (
	reDigit = regexp.MustCompile(`[\d]`)
	steps = []string{"O", "|", "/", "\\", "/", "\\"}
	baseBoard = `
 +----+
 |    |
 1    |
324   |
5 6   |
	  |
==========`
)

func NewGame(hgc HangmanContract) *HangmanGame {
	return &hgc{}
}

func (hg *HangmanGame) LoadBases() {
	// TODO: created board sort word...
}