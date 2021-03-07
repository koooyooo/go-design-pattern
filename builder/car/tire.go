package car

type (
	Tire interface {
		Grip() string
	}
	SportsTire   struct{}
	NormalTire   struct{}
	StudlessTire struct{}
)

func (t SportsTire) Grip() string {
	return "Grip(Sport)"
}
func (t NormalTire) Grip() string {
	return "Grip(Normal)"
}
func (t StudlessTire) Grip() string {
	return "Grip(Studless)"
}
