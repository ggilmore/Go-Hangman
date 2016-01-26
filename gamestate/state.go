package gamestate

import (
	"../messages/hangman"
	"../utilities"
)

/*
Represents a game of hangman. Keeps track of what the target word is (targetWord),
what letters in the word that have been spelled (wordLetterStatus),
every letter that has been guessed so far(guessedLetters),
and the number of guesses the player has remaining (triesRemaining)

*/

type GameState struct {
	targetWord       string
	wordLetterStatus utilities.Set
	guessedLetters   utilities.Set
	triesRemaining   int
}

func New(targetWord string, maxTries int) GameState {
	wordLetterStatus := make(utilities.Set)
	for _, l := range targetWord {
		wordLetterStatus[l] = false
	}

	return GameState{targetWord, wordLetterStatus, make(utilities.Set), maxTries}
}

func (state *GameState) checkWordSpelled() bool {
	for _, l := range state.targetWord {
		if guessed, _ := state.wordLetterStatus[l]; !guessed {
			return false
		}
	}
	return true
}

func (state *GameState) checkPlayerLost() bool {
	return !state.checkWordSpelled() && (state.triesRemaining <= 0)
}

func (state *GameState) checkPlayerWon() bool {
	return state.checkWordSpelled() && (state.triesRemaining > 0)
}

func (state *GameState) generateWordStatusRepresentation() string {

	out := make([]rune, len(state.targetWord))

	for i, letter := range state.targetWord {
		if guessed, _ := state.wordLetterStatus[letter]; guessed {
			out[i] = letter
		} else {
			out[i] = '_'
		}
	}

	return utilities.Join(out)
}

func (state *GameState) checkGuess(letter rune) bool {
	_, validLetter := state.wordLetterStatus[letter]

	if validLetter {
		return true
	} else {
		return false
	}

}

func (state *GameState) Guess(guess rune) hangman.HangmanMessage {

	state.guessedLetters[guess] = true

	if state.checkGuess(guess) {
		state.wordLetterStatus[guess] = true
	} else {
		state.triesRemaining = state.triesRemaining - 1
	}

	guessedLettersCopy := utilities.SetCopy(state.guessedLetters)

	Message := hangman.DefaultHangmanMessage{
		state.triesRemaining, state.generateWordStatusRepresentation(),
		guessedLettersCopy, state.targetWord}

	if state.checkPlayerLost() {
		return hangman.YouLost{Message}
	} else if state.checkPlayerWon() {
		return hangman.YouWon{Message}
	} else {
		return hangman.StillPlaying{Message}
	}
}
