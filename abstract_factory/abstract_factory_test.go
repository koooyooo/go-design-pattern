// [目的]
// Abstract Factoryの目的は 複雑な構成を組み上げることです。
//
// [概要]
package abstract_factory_pattern

import (
	"testing"

	"github.com/koooyooo/go-design-pattern/abstract_factory/car_factory"
	"github.com/stretchr/testify/assert"
)

func TestAbstractFactorySports(t *testing.T) {
	factory, err := car_factory.NewSportsCarFactory()
	assert.NoError(t, err)

	car, err := factory.CreateCar()
	assert.NoError(t, err)

	runMsg, err := car.Run()
	assert.NoError(t, err)

	assert.Equal(t, "sustain(sports),start(sports),grip(sports),grip(sports),grip(sports),grip(sports)", runMsg)
}

func TestAbstractFactoryFamily(t *testing.T) {
	factory, err := car_factory.NewFamilyCarFactory()
	assert.NoError(t, err)

	car, err := factory.CreateCar()
	assert.NoError(t, err)

	runMsg, err := car.Run()
	assert.NoError(t, err)

	assert.Equal(t, "sustain(family),start(family),grip(family),grip(family),grip(family),grip(family)", runMsg)
}
