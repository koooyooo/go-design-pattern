// [目的]
// Abstract Factoryの目的は 複雑な製品を組み上げることです。
// 通常のFactoryでは単純な製品しか生成できませんが、AbstractFactoryは複数の内部Factoryを連携させることでより複雑な製品の生成が可能です。
//
// [概要]
// AbstractFactoryの特徴は 複数の内部Factoryの調整を行う TemplateMethodとして振る舞う点です。
// DIコンテナの様に複数のインターフェイスが依存する複雑な生成物を用意するには、生成側のFactoryも構造化させる必要があります。
// そこで、AbstractFactoryが以下の調整を行う訳です。
// - FactoryMethodの一覧を宣言し、必要な実装を宣言
// - FactoryMethodの実装を活用し、それらの成果物を組み合わせた複雑な製品を製造 (するFactoryを返却)
//
// こうすることで、利用側は AbstractFactoryに依頼するだけで、複雑な製品を取得することができますし、
// 実装側は AbstractFactoryの指示に従ったFactoryMethodの実装を提供するだけで、組み合わせ等は気にしなくて良くなる訳です
//
// [シナリオ]
// 今回のシナリオは車の工場です。
// 同じ AbstractFactoryに、スポーツカーを組み立てる実装を注入したものと、ファミリーカーを組み立てる実装を注入したものを用意しました。
// 2系統の実装が存在しますが、複雑性を軽減した差分の実装を提供するだけで新規の成果物を生成することが可能です。
// 部品を生成するロジックは別々ですが、基本的な製造ラインは同じものを使っている訳ですから。
package abstract_factory_pattern

import (
	"testing"

	"github.com/koooyooo/go-design-pattern/abstract_factory/car_factory"
	"github.com/stretchr/testify/assert"
)

// スポーツカーの工場による実装です。
// テストコードでは分かりづらいですが、NewSportsCarFactory()で返る工場が AbstractFactoryです。
// AbstractFactoryの実装は全てスポーツカーの部品を製造するFactoryMethodで組み上げています。
func TestAbstractFactorySports(t *testing.T) {
	// スポーツカーの工場を作成
	factory, err := car_factory.NewSportsCarFactory()
	assert.NoError(t, err)

	// 工場にスポーツカーの製造を依頼
	car, err := factory.CreateCar()
	assert.NoError(t, err)

	// スポーツカーを走らせ、出力が正しいかを確認
	runLog, err := car.Run()
	assert.NoError(t, err)
	assert.Equal(t, "sustain(sports),start(sports),grip(sports),grip(sports),grip(sports),grip(sports)", runLog)
}

// ファミリーカーの工場による実装です。
// テストコードでは分かりづらいですが、NewFamilyCarFactory()で返る工場が AbstractFactoryです。
// AbstractFactoryの実装は全てファミリーカーの部品を製造するFactoryMethodで組み上げています。
func TestAbstractFactoryFamily(t *testing.T) {
	// ファミリーカーの工場を作成
	factory, err := car_factory.NewFamilyCarFactory()
	assert.NoError(t, err)

	// 工場にファミリーカーの製造を依頼
	car, err := factory.CreateCar()
	assert.NoError(t, err)

	// ファミリーカーを走らせ、出力が正しいかを確認
	runLog, err := car.Run()
	assert.NoError(t, err)
	assert.Equal(t, "sustain(family),start(family),grip(family),grip(family),grip(family),grip(family)", runLog)
}
