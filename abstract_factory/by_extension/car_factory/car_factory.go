package car_factory

type (
	// AbstractCarFactory は抽象Factoryです。
	// 本来は継承を用いて TemplateMethod的に生成フローを制御します。
	// Golangの場合、継承の代替としてEmbedを使うと親に子を代入できません。そのため、まず代入用のinterfaceを用意します。
	AbstractCarFactory interface {
		CreateCar() (*Car, error)
	}

	// abstractCarFactoryImpl はAbstractCarFactoryのベースとなる実装です。
	// 本来のAbstractFactoryは継承なので、①親の型への代入・②テンプレートフローの提供 の双方が行なえますが、
	// Golangで①を interfaceで実現した場合、②のテンプレートメソッドとなる関数は実装できないので、基底型を用意する必要があります
	abstractCarFactoryImpl struct {
		// 製造過程の関数群も、実際は子クラスでのOverrideを用いますが、interfaceでの切り替えで実現します。
		carFactory
	}
)

// CreateCar は車を生成します
func (f abstractCarFactoryImpl) CreateCar() (*Car, error) {
	// bodyを用意
	b, err := f.createBody()
	if err != nil {
		return nil, err
	}
	// engineを搭載
	e, err := f.createEngine()
	if err != nil {
		return nil, err
	}
	// tireを取り付け
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

type carFactory interface {
	createBody() (body, error)
	createEngine() (engine, error)
	createTire() (tire, error)
}
