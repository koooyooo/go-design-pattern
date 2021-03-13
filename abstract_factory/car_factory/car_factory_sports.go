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

func NewSportsCarFactory() (*AbstractCarFactory, error) {
	return &AbstractCarFactory{
		createBody: func() (body, error) {
			return &sportsBody{}, nil
		},
		createEngine: func() (engine, error) {
			return &sportsEngine{}, nil
		},
		createTire: func() (tire, error) {
			return &sportsTire{}, nil
		},
	}, nil
}
