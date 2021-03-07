package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/builder/car"
)

func TestBuilder(t *testing.T) {
	b := car.NewBuilder()
	car, err := b.InstallEngine(true).InstallHandle(true).InstallSportTire().Build()
	assert.NoError(t, err)
	assert.Equal(t, "Drive(V8),Control(Sport),Grip(Sport),Grip(Sport),Grip(Sport),Grip(Sport)", car.Drive())
}
