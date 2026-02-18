package animations

type Animation struct {
	First        int
	Last         int
	Step         int
	Speed        float32
	FrameCounter float32
	Frame        int
}

func NewAnimation(first, last, step int, speed float32) *Animation {
	return &Animation{
		First:        first,
		Last:         last,
		Step:         step,
		Speed:        speed,
		FrameCounter: speed,
		Frame:        first,
	}
}

func (a *Animation) Update() {
	a.FrameCounter -= 1.0
	if a.FrameCounter < 0.0 {
		a.FrameCounter = a.Speed
		a.Frame += a.Step
		if a.Frame > a.Last {
			a.Frame = a.First
		}
	}
}

func (a *Animation) GetFrame() int {
	return a.Frame
}
