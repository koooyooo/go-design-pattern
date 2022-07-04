package extension

// Mammal は哺乳類を表現します
type Mammal struct {
}

func (m Mammal) Do() string {
	return "do something"
}

// Human は人類を表現します
type Human struct {
	Mammal
}

// Do で定義をOverrideします
func (h Human) Do() string {
	return h.Mammal.Do() + " by human"
}
