package components

type Gravity struct {
	Acceleration     float64
	TerminalVelocity float64
}

func NewGravity() Gravity {
	return Gravity{
		Acceleration:     0.5,
		TerminalVelocity: 10,
	}
}
