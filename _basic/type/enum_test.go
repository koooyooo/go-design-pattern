package _type

import (
	"github.com/koooyooo/go-design-pattern/_basic/type/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnum(t *testing.T) {
	assert.Equal(t, enum.Hello, enum.Hello)
	assert.NotEqual(t, enum.Hello, enum.World)

	assert.NotEqual(t, 1, enum.Hello)
	assert.Equal(t, 1, int(enum.Hello))
	assert.Equal(t, "Hello", enum.Hello.String())
}

func TestStructEnum(t *testing.T) {
	assert.Equal(t, enum.January, enum.January)
	assert.NotEqual(t, enum.January, enum.February)

	assert.Equal(t, "[1] January", enum.January.String())
}
