package manager

import (
	"bufio"
	"github.com/ggilmore/Go-Hangman/gamestate"
	"github.com/ggilmore/Go-Hangman/messages/managermsgs"
	"net"
)

type manager struct {
	game gamestate.GameState

	conn net.Conn

	wordManager *WordManager
}

func (m *manager) Serve() {
	readWrite := bufio.NewReadWriter(bufio.NewReader(m.conn), bufio.NewWriter(m.conn))

	for {
		input, err := readWrite.Reader.ReadString('\n')
		if err != nil {
			break
		}
		playerMessage, parseErr := ParsePlayerInput(input)

		if parseErr != nil {
			readWrite.Writer.WriteString(parseErr.Error())

		} else {
			switch playerMessage.(type) {
			case managermsgs.Peek:

			}
		}

	}
}
