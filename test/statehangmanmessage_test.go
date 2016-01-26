package test

import (
	"testing"

	"../gamestate"
	"../messages/hangman"
	"../utilities"
	"reflect"
	"strings"
)

func blankState() gamestate.GameState {
	return gamestate.New("", 10)
}

func bookState() gamestate.GameState {
	return gamestate.New("book", 10)
}

func trimType(t string) string {
	return strings.TrimLeft(t, ".")
}


/*
getting the types so that we can do the comparisons in the test
 */
var (
	set = make(utilities.Set)
	message = hangman.DefaultHangmanMessage{0, "___", &set, "dummy"}

	YouLostType = reflect.TypeOf(hangman.YouLost{message}).Name()
	YouWonType = reflect.TypeOf(hangman.YouWon{message}).Name()
	StillPlayingType = reflect.TypeOf(hangman.StillPlaying{message}).Name()
)


type testStructure struct {
	testState   gamestate.GameState
	input       string
	desiredOuts []string
}

var tests = []testStructure{
	testStructure{
		bookState(),
		"book",
		[]string{StillPlayingType, StillPlayingType, StillPlayingType, YouWonType}}}



func TestWordSpelledBlank(t *testing.T) {
	for _, test := range tests{
		state := test.testState
		for i, letter := range test.input {
			response := state.Guess(letter)
			if trimType(reflect.TypeOf(response).Name()) != trimType(test.desiredOuts[i]){
				t.Error("Expected Type:", test.desiredOuts[i], "Got: ", trimType(reflect.TypeOf(response).Name()))
			}
		}
	}
}