package animals

type Dog struct {
	Name string
}

func (d Dog) MakeNoise() string {
	return "Wuff! Wuff!"
}
