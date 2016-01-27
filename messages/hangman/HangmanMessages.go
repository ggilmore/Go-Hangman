package hangman

import utilities "../../utilities"

type HangmanMessage interface {
	RemainingTries() int
	WordRepr() string
	LettersTriedSet() *utilities.RuneSet
	Word() string
}

type DefaultHangmanMessage struct {
	Tries              int
	WordRepresentation string
	LettersTried      *utilities.RuneSet
	TargetWord         string
}

func (message DefaultHangmanMessage) RemainingTries() int {
	return message.Tries
}

func (message DefaultHangmanMessage) WordRepr() string {
	return message.WordRepresentation
}

func (message DefaultHangmanMessage) LettersTriedSet() *utilities.RuneSet {
	return message.LettersTried
}

func (message DefaultHangmanMessage) Word() string {
	return message.TargetWord
}

type StillPlaying struct {
	DefaultHangmanMessage
}

type YouWon struct {
	DefaultHangmanMessage
}

type YouLost struct {
	DefaultHangmanMessage
}







