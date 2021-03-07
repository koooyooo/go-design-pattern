package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/builder/car"
)

func TestBuilder(t *testing.T) {
	// ビルダーを生成
	b := car.NewBuilder()

	// ビルダーを利用して複雑なインスタンスを構築
	// 単一関数で実現でも問題ないが、敢えて段階的にビルドする形式にした
	car, err := b.InstallEngine(true).InstallHandle(true).InstallSportTire().Build()

	// 無事にビルドされていることを確認
	assert.NoError(t, err)
	assert.Equal(t, "Drive(V8),Control(Sport),Grip(Sport),Grip(Sport),Grip(Sport),Grip(Sport)", car.Drive())
}
