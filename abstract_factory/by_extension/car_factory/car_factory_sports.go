package car_factory

type (
	sportsBody   struct{}
	sportsEngine struct{}
	sportsTire   struct{}
)

func (sb sportsBody) sustain() string {
	return "sustain(sports)"
}

func (sb sportsEngine) start() string {
	return "start(sports)"
}

func (sb sportsTire) grip() string {
	return "grip(sports)"
}

type sportsCarFactory struct{}

func (f *sportsCarFactory) createBody() (body, error) {
	return &sportsBody{}, nil
}

func (f *sportsCarFactory) createEngine() (engine, error) {
	return &sportsEngine{}, nil
}

func (f *sportsCarFactory) createTire() (tire, error) {
	return &sportsTire{}, nil
}

func NewSportsCarFactory() (AbstractCarFactory, error) {
	return &abstractCarFactoryImpl{&sportsCarFactory{}}, nil
}
