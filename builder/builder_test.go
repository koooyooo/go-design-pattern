// Builder は複雑なインスタンスの生成を隠蔽するパターンです。
// Build プロセスにおける調停役(Mediator)であり、フロント(Facade)でもあります。
// ある程度の要件を職人に伝えれば、職人が製品を作ってくれる、その職人がBuilderとなります。
package builder_pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/builder/car"
)

// ここでは車のビルダーを用意します。
// 車の製造は大変ですが、エンジン、ハンドル、タイヤの仕様を伝えるだけで、職人が作成してくれる状況をイメージしています。
func TestBuilder(t *testing.T) {

	// 車のビルダーを生成
	b := car.NewBuilder()

	// ビルダーを利用して複雑なインスタンスを構築
	// 単一関数で実現でも問題ないが、敢えて段階的にビルドする形式に
	car, err := b.InstallEngine(true).InstallHandle(true).InstallSportTire().Build()

	// 無事にビルドされていることを確認
	assert.NoError(t, err)
	assert.Equal(t, "Drive(V8),Control(Sport),Grip(Sport),Grip(Sport),Grip(Sport),Grip(Sport)", car.Drive())
}
