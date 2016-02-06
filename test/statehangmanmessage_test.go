package test

import (
	"testing"

	"../gamestate"
	"../messages/hangman"
	"../utilities"
	"reflect"
	"strings"
)

//fixture
func bookState1Try() gamestate.GameState {
	state, _ := gamestate.New("book", 1, false)
	return state
}

//helper function to clean up the types that we get back from reflection
func trimType(t string) string {
	return strings.TrimLeft(t, ".")
}

/*
getting the types so that we can do the comparisons in the test
*/
var (
	set     = make(utilities.RuneSet)
	message = hangman.DefaultHangmanMessage{0, "___", &set, "dummy"}

	YouLostType      = reflect.TypeOf(hangman.YouLost{message}).Name()
	YouWonType       = reflect.TypeOf(hangman.YouWon{message}).Name()
	StillPlayingType = reflect.TypeOf(hangman.StillPlaying{message}).Name()
)

type testStruct struct {
	descr              string              //description of the test
	testState          gamestate.GameState //game state that we're testing
	input              string              //the series of input characters to the gamestate
	wordRep            []string            // the representation of the word after each input
	tryCounts          []int               //the number of tries remaining after each input
	desiredOutMessages []string            //the types of the output messages that we expect
}

//battery of tests to run
var tests = []testStruct{
	testStruct{
		"Target word: book, input: 'book', should recieve valid sequence of hangman messages and win at the end",
		bookState1Try(),
		"book",
		[]string{"b___", "boo_", "boo_", "book"},
		[]int{1, 1, 1, 1},
		[]string{StillPlayingType, StillPlayingType, StillPlayingType, YouWonType}},

	testStruct{
		"Target word: book, input: 'bok', should recieve valid sequence of hangman messages and win at the end",
		bookState1Try(),
		"bok",
		[]string{"b___", "boo_", "book"},
		[]int{1, 1, 1},
		[]string{StillPlayingType, StillPlayingType, YouWonType}},

	testStruct{
		"Target word: book, input: 'boo', should still be playing at the end",
		bookState1Try(),
		"boo",
		[]string{"b___", "boo_", "boo_"},
		[]int{1, 1, 1},
		[]string{StillPlayingType, StillPlayingType, StillPlayingType}},

	testStruct{
		"Target word: book, input: 'c', should loose at the end",
		bookState1Try(),
		"c",
		[]string{"____"},
		[]int{0},
		[]string{YouLostType}}}

func TestGameStateGuessSequence(t *testing.T) {
	for _, test := range tests {
		t.Log(test.descr)
		inputSoFar := make([]rune, len(test.input))
		for i, letter := range test.input {
			inputSoFar[i] = letter

			response := test.testState.Guess(letter)
			responseType := trimType(reflect.TypeOf(response).Name())

			desiredType := trimType(test.desiredOutMessages[i])

			if responseType != desiredType {
				t.Error("Input So Far:", utilities.JoinRunes(inputSoFar), "Expected Type:", desiredType, "Got: ", responseType)
			} else if response.WordRepr() != test.wordRep[i] {
				t.Error("Input So Far:", utilities.JoinRunes(inputSoFar), "Expected Word Rep:", test.wordRep[i], "Got: ", response.WordRepr())
			} else if response.RemainingTries() != test.tryCounts[i] {
				t.Error("Input So Far:", utilities.JoinRunes(inputSoFar), "Expected # Tries:", test.tryCounts[i], "Got: ", response.RemainingTries())
			}

		}
	}
}
