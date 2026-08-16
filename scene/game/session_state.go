package game

// TODO: manage game session. Lives, score, asteriods waves, etc
type SessionState struct {
	lives     int
	WaveCount int
}

func NewSessionState() SessionState {
	return SessionState{
		lives:     3,
		WaveCount: 1,
	}
}

func (s *SessionState) Update(frame Frame) {}
