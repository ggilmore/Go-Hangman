package manager

import (
	"../messages/managermsgs"
	"errors"
	"github.com/ggilmore/Go-Hangman/utilities"
	"strings"
)

/**
Parses player input, performs basic validation, and returns the relevant PlayerInputMessage types
*/

const (
	PEEK_CMD    = "peek"
	SETUP_CMD   = "setup"
	GUESS_CMD   = "guess"
	DSC_CMD     = "disconnect"
	RESET_CMD   = "reset"
	CONNECT_CMD = "connect"

	//Errors
	ALPHANUMERIC_CHARS_ONLY_ERROR = "Only alphabetical runes allowed"
	USAGE                         = PEEK_CMD + "\n" + RESET_CMD + "\n" + DSC_CMD + "\n" + GUESS_CMD + "[guessed_letter] \n" +
		CONNECT_CMD + "[game_id] \n" + SETUP_CMD + "[target_word] \n"
)

func ParsePlayerInput(input string) (managermsgs.PlayerInputMessage, error) {
	cleanInput := strings.ToLower(strings.TrimSpace(input))

	var message managermsgs.PlayerInputMessage
	var err error

	//check for special characters (illegal)
	for _, l := range cleanInput {
		if isValid, _ := utilities.SET_OF_LETTERS[l]; !isValid {
			return message, errors.New(ALPHANUMERIC_CHARS_ONLY_ERROR)
		}
	}

	args := strings.Split(cleanInput, " ")

	switch len(args) {
	case 1:
		switch args[0] {
		case PEEK_CMD:
			message = managermsgs.Peek{""}
		case RESET_CMD:
			message = managermsgs.Reset{""}
		case DSC_CMD:
			message = managermsgs.Disconnect{""}
		default:
			err = errors.New(USAGE)
		}
	case 2:
		switch args[0] {
		case GUESS_CMD:
			message = managermsgs.GuessLetter{args[1]}
		case CONNECT_CMD:
			message = managermsgs.Connect{args[1]}
		case SETUP_CMD:
			message = managermsgs.Setup{args[1]}
		default:
			err = errors.New(USAGE)
		}
	default:
		err = errors.New(USAGE)
	}

	return message, err
}
