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
	return &AbstractCarFactoryImpl{&familyCarFactory{}}, nil

	//// 抽象の工場に実装を注入します。
	//// 本来は継承を用いてMethodのOverrideで差分実装しますが、Golangには継承が無いので関数を注入し委譲することでOverrideします。
	//return &AbstractCarFactory{
	//	createBody: func() (body, error) {
	//		return &familyBody{}, nil
	//	},
	//	createEngine: func() (engine, error) {
	//		return &familyEngine{}, nil
	//	},
	//	createTire: func() (tire, error) {
	//		return &familyTire{}, nil
	//	},
	//}, nil
}
