package factory_method

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/factory_method/factory"
)

func TestFactoryMethod(t *testing.T) {
	// Factory経由で Storageを取得するがその実体は知らない（疎結合）
	s := factory.GetStorage()
	err := s.Store([]byte("Hello"))
	assert.NoError(t, err)
}
