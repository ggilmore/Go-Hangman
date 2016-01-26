package hangman

import utilities "../../utilities"

type HangmanMessage interface {
	RemainingTries() int
	WordRepr() string
	LettersTriedSet() *utilities.Set
	Word() string
}

type DefaultHangmanMessage struct {
	Tries              int
	WordRepresentation string
	LettersTried      *utilities.Set
	TargetWord         string
}

func (message DefaultHangmanMessage) RemainingTries() int {
	return message.Tries
}

func (message DefaultHangmanMessage) WordRepr() string {
	return message.WordRepresentation
}

func (message DefaultHangmanMessage) LettersTriedSet() *utilities.Set {
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







