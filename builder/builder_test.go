// [目的]
// Builder はインスタンス構築時の複雑性を隠蔽するパターンです。
//
// [概要]
// 小難しく複雑な製品がほしい時、自分で作るより要件だけ伝えて職人に作って貰いたいですよね、その職人がBuilderとなります。
// 複雑性を隠蔽するパターンは 調停役(Mediator)や、窓口(Facade) 等複数のパターンがありますが、Builderは構築段階におけるものです。
// 特徴として、特段ポリモーフィズムや継承といったテクニックは使いません。
// 複雑性さえ隠蔽できれば立派なビルダーなのですが、構築段階ごとに自身のインスタンスを返すことで
// `Builder.setXX().setXX().build()` と連鎖的にコール可能な形式で実現されることが多いです。
//
// [実装例]
// 文字列操作における Javaの StringBuilder(), Golangの bytes.Buffer
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
