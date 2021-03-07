package car

import "errors"

type builder struct {
	car *car
}

func NewBuilder() *builder {
	return &builder{
		car: &car{},
	}
}

func (b *builder) InstallEngine(highPowered bool) *builder {
	if highPowered {
		b.car.engine = &V8Engine{}
	} else {
		b.car.engine = &V4Engine{}
	}
	return b
}

func (b *builder) InstallHandle(isSport bool) *builder {
	if isSport {
		b.car.handle = SportHandle{}
	} else {
		b.car.handle = NormalHandle{}
	}
	return b
}

func (b *builder) InstallSportTire() *builder {
	b.car.tires = []Tire{
		&SportsTire{},
		&SportsTire{},
		&SportsTire{},
		&SportsTire{},
	}
	return b
}

func (b *builder) BuildNormalTire() *builder {
	b.car.tires = []Tire{
		&NormalTire{},
		&NormalTire{},
		&NormalTire{},
		&NormalTire{},
	}
	return b
}

func (b *builder) BuildStudlessTire() *builder {
	b.car.tires = []Tire{
		&StudlessTire{},
		&StudlessTire{},
		&StudlessTire{},
		&StudlessTire{},
	}
	return b
}

func (b *builder) Build() (*car, error) {
	if b.car.engine == nil {
		return nil, errors.New("no engine")
	}
	if b.car.handle == nil {
		return nil, errors.New("no handle")
	}
	if len(b.car.tires) != 4 {
		return nil, errors.New("not 4 tires")
	}
	return b.car, nil
}
