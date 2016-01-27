package managermsgs


type PlayerInputMessage interface{
	Payload() string
}

type PlayerInputMessageDefault struct {
	payload string
}

func (m PlayerInputMessageDefault) Payload() string {
	return m.payload
}

type Connect struct {
	PlayerInputMessageDefault
}

type Peek struct {
	PlayerInputMessageDefault
}

type Setup struct {
	PlayerInputMessageDefault
}

type GuessLetter struct {
	PlayerInputMessageDefault
}

type Disconnect struct {
	PlayerInputMessageDefault
}

type Reset struct {
	PlayerInputMessageDefault
}

