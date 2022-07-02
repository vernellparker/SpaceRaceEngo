package messages

type ScoreMessage struct {
	Player int
}

func (msg *ScoreMessage) Type() string {
	return "ScoreMessage"
}
