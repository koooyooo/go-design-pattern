// [目的]
// Singleton は唯一のインスタンスを生成するパターンです。
// Singleton の特徴はクラス(struct)やインターフェイス同士の関連ではなく、クラス内閉じられたパターンであるという点です。
//
// Singleton の目的は複数存在します。
// 1. 単一のインスタンスの提示による状態共有を実現するため (ex. Observer.getInstance().addAccessCount(1))
// 2. Enumの様に種別を表現するクラスの種別数とインスタンス数を一致させ正確に表現するため (同値性だけでなく同一性まで担保可能)
// 3. Flyweightパターンの様にインスタンス生成の抑制を通じて性能を担保する仕組みを提供するため
// 4. Flyweightと似ているが、Serviceの様に 操作中心で状態を持たないインスタンスを組み合わせて提供するため (DI)
//
// [概要]
// 先に述べたように、様々な目的で活用可能なのが Singletonですが、自分で実装する機会は少ないかも知れません。
// 理由は言語機能的に、またはF/W的に実現されることが多いパターンであるからです。
// 例えば Enumは Singletonの独自実装が可能なものの、言語がサポートしている場合が多いです。
// また、DIコンテナの様に内部では単一インスタンス管理をしているものの、利用側は意識せずに利用しているはずです。
//
// とは言え、生成インスタンスの数を制限する、特にSingletonの様に単一に絞る…というテクニックは一度覚えておくと様々な局面で利用可能
// となりますので、身につけておくと良いと思います。
//
// [実現]
// 以下の仕組みでこれを実現します。
// - 外部からインスタンスを生成できなくする
// - Javaではコンストラクタを`private`にすることで、Goではstructを(小文字で開始し)`package private`にすることで実現する
// - インスタンス取得用の関数を別途用意し、そこでは常に同じインスタンスを返す様に実装する
package singleton_pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/singleton/singleton"
)

// Singletonは唯一のインスタンスを用意し、それ以外を生成できなくするパターンです。
// そのため、唯一のインスタンスしか存在しないことを保証できます。
func TestSingleton(t *testing.T) {
	// Instance()以外で生成できない => singleton.singleton{} は不可能
	s1 := singleton.Instance()
	s2 := singleton.Instance()

	// インスタンスの内容が同じだけでなく
	assert.Equal(t, s1, s2)
	// インスタンスの実体も同じ
	assert.True(t, s1 == s2)
}

// 逆に一般的な生成パターンは、コンストラクタ関数の中でインスタンスを生成します。
// そのため、同じ値を持ちますが(同値)、参照は同じ(同一)ではありません。
func TestNonSingleton(t *testing.T) {
	s1 := singleton.NewNormal()
	s2 := singleton.NewNormal()

	// インスタンスの内容は同じだが
	assert.Equal(t, s1, s2)
	// インスタンスの実体は異なる
	assert.False(t, s1 == s2)
}
