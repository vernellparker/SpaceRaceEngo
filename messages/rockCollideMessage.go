package messages

type RockCollideMessage struct {
	PlayerId int
}

func (RockCollideMessage) Type() string {
	return "RockCollideMessage"
}
