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

	"github.com/koooyooo/go-design-pattern/builder/pc"
)

func TestBuilder(t *testing.T) {
	builder := pc.NewPCBuilder()
	builtPC := builder.SetUpBaseUnit(550).SetCPU(2).SetMemory(16).SetSSD(256).Build()

	spec := builtPC.Spec()
	assert.Equal(t, "Power: 550, CPU: 2, Mem: 16, SSD: 256", spec)
}
