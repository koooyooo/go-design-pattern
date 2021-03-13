package car_factory

type (
	// AbstractCarFactory は抽象Factoryです
	// 本来は継承を用いて TemplateMethod的に生成フローを制御しますが
	// Golangは継承がないので委譲で表現します。具体的には抽象メソッド部分を関数にした通常の structで表現します
	AbstractCarFactory struct {
		createBody
		createEngine
		createTire
	}

	// 抽象メソッドの代替として用意した関数群
	createBody   func() (body, error)
	createEngine func() (engine, error)
	createTire   func() (tire, error)
)

// CreateCar は車を生成します
func (f AbstractCarFactory) CreateCar() (*Car, error) {
	// bodyを作成し
	b, err := f.createBody()
	if err != nil {
		return nil, err
	}
	// engineを載せ
	e, err := f.createEngine()
	if err != nil {
		return nil, err
	}
	// tireを搭載します
	var tires []tire
	for i := 0; i < 4; i++ {
		t, err := f.createTire()
		tires = append(tires, t)
		if err != nil {
			return nil, err
		}
	}
	return &Car{b, e, tires}, nil
}
