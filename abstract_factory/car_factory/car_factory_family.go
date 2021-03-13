package car_factory

type (
	familyBody   struct{}
	familyEngine struct{}
	familyTire   struct{}
)

func (sb familyBody) sustain() string {
	return "sustain(family)"
}

func (sb familyEngine) start() string {
	return "start(family)"
}

func (sb familyTire) grip() string {
	return "grip(family)"
}

// FamilyCarFactory
func NewFamilyCarFactory() (*AbstractCarFactory, error) {
	return &AbstractCarFactory{
		createBody: func() (body, error) {
			return &familyBody{}, nil
		},
		createEngine: func() (engine, error) {
			return &familyEngine{}, nil
		},
		createTire: func() (tire, error) {
			return &familyTire{}, nil
		},
	}, nil
}
