package game

// TODO: manage game session. Lives, score, asteriods waves, etc
type SessionState struct{}

func NewSessionState() SessionState {
	return SessionState{}
}

func (s *SessionState) Update(frame Frame) {}
