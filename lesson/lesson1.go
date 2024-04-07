package solution

type Voicer interface {
	Voice() string
}

type Cat struct {
	// …
}

type Cow struct {
	// …
}

type Dog struct {
	// …
}

func (c Cat) Voice() {
	return "Мяу"
}

func (c Cow) Vioce() {
	return "Мууу"
}

func (c Dog) Vioce() {
	return "Гав"
}
