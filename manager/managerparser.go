package manager

import (
	"strings"
	"../messages/managermsgs"
	"errors"
	"github.com/ggilmore/hangman/utilities"
)


const (
	PEEK_CMD = "peek"
	SETUP_CMD = "setup"
	GUESS_CMD = "guess"
	DSC_CMD = "disconnect"
	RESET_CMD = "reset"
	CONNECT_CMD = "connect"
)

const (
	ALPHA_CHARS_ONLY_ERROR = "Only alphabetical runes allowed"
	SETUP_
)


func ParsePlayerInput(input string) (managermsgs.PlayerInputMessage, error) {
	cleanInput := strings.ToLower(strings.TrimSpace(input))

	var message managermsgs.PlayerInputMessage

	//check for special characters (illegal)
	for _, l := range cleanInput {
		if isLetter, _ := utilities.SETOFLETTERS[l]; !isLetter {
			return message, errors.New(ALPHA_CHARS_ONLY_ERROR)
		}
	}

	args := strings.Split(cleanInput, " ")

	switch len(args){
	case 1:
		switch args[0] {
		case PEEK_CMD:
			return managermsgs.Peek{""}, nil
		case RESET_CMD:
			return managermsgs.Reset{""}, nil
		case DSC_CMD:
			return managermsgs.Disconnect{""}, nil
		default:
			return message
		}
	case 2:
		switch args[0] {
		case GUESS_CMD:
			return managermsgs.GuessLetter{args[1]}, nil
		case CONNECT_CMD:
			return managermsgs.Connect{args[1]}, nil
		case SETUP_CMD:
			return managermsgs.Setup{args[1]}, nil
		default:
			return
		}
	default:
		return
	}





}