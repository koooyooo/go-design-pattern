package car_factory

import "fmt"

type (
	// AbstractCarFactory は抽象Factoryです。
	// 本来は継承を用いて TemplateMethod的に生成フローを制御しますが、Golangは継承がないので委譲で表現します。
	//具体的には抽象メソッド部分を関数にした通常の structで表現します

	// AbstractCarFactory は抽象Factoryです。
	// 本来は継承を用いて TemplateMethod的に生成フローを制御します。
	// Golangの場合、継承の代替としてEmbedを使うと親に子を代入できません。そのため、まず代入用のinterfaceを用意します。
	AbstractCarFactory interface {
		CreateCar() (*Car, error)
	}

	// AbstractCarFactoryImpl はAbstractCarFactoryのベースとなる実装です。
	// 本来のAbstractFactoryは継承なので、①親の型への代入・②テンプレートフローの提供 の双方が行なえますが、
	// Golangで①を interfaceで実現した場合、②のテンプレートメソッドとなる関数は実装できないので、既定型を用意する必要があります
	AbstractCarFactoryImpl struct {
		// テンプレートメソッドが操作する関数群
		carFactory
	}

	carFactory interface {
		createBody() (body, error)
		createEngine() (engine, error)
		createTire() (tire, error)
	}

	//// 抽象メソッドの代替として用意した関数群
	//createBody   func() (body, error)
	//createEngine func() (engine, error)
	//createTire   func() (tire, error)
)

// CreateCar は車を生成します
func (f AbstractCarFactoryImpl) CreateCar() (*Car, error) {
	fmt.Println("F:", f)
	fmt.Println("F-cf:", f.carFactory)
	fmt.Println("F-cb", f.createBody)
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
