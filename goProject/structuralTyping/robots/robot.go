package robots

type Robot struct {
	Model string
}

func (r Robot) MakeNoise() string {
	return "Beep Beep. Initializing starting ..."
}
