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

type familyCarFactory struct{}

func (f *familyCarFactory) createBody() (body, error) {
	return &familyBody{}, nil
}

func (f *familyCarFactory) createEngine() (engine, error) {
	return &familyEngine{}, nil
}

func (f *familyCarFactory) createTire() (tire, error) {
	return &familyTire{}, nil
}

func NewFamilyCarFactory() (AbstractCarFactory, error) {
	return &abstractCarFactoryImpl{&familyCarFactory{}}, nil
}
