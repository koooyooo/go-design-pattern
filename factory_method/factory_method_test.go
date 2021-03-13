// [目的]
// FactoryMethod はインスタンスの生成に関するパターンです。
// このパターンの目的は、ポリモーフィズムにおけるインスタンス生成・インターフェイス代入の局面を隠し、疎結合を完成させることです。
//
// [概要]
// ポリモーフィズムとは仕事を決めてやり方を決めないことです。実際は、インターフェイスが仕事を決め、それを実現するクラスがやり方を決めます。
// インターフェイスには任意のクラスを適用できますので、やり方の違うクラスを適用すれば、様々なやり方で仕事を全うすることができます。
// また、利用側のクラスにはインターフェイスしか見えませんので、適用されたクラスが変わったとしても、クラスの修正や再テストは不要です。
//
// しかし、インターフェイスを利用するにはクラスを生成しインターフェイス型に代入しなければなりません。
// この生成・代入の瞬間を利用側のクラスが知ってしまうと、疎結合が崩れ、クラスの修正や再テストが必要になってしまいます。
// そこで、利用側のクラスは FactoryMethodにこの仕事を依頼します。 そうすることで、利用側は実現するクラスに対して何も知らずに済むのです。
//
// [実例・類似例]
// - DIコンテナは Factoryの完成形と言えるべきものです。単純な隠蔽だけでなく、関連を完成させ、Factory自体も見せないといった付加価値まで提供しています。
// - Singletonの getInstance() はインスタンスの実体ではなくインスタンスの生成を隠蔽しています。目的は異なりますが作りは似ています。
package factory_method

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/factory_method/factory"
)

// FactoryMethodを適用した場合
// Storageの実体を知ることなく利用できる (疎結合)
func TestFactoryMethod(t *testing.T) {
	// Factory経由で実体を隠蔽
	s := factory.GetStorage()
	err := s.Store([]byte("Hello"))
	assert.NoError(t, err)
}

// FactoryMethodを適用しない場合
// Storageの実体を知ってしまう (密結合)
func TestNonFactory(t *testing.T) {
	var s factory.Storage
	// 生成・代入プロセスで実体と密結合
	s = &factory.NopeStorage{}
	err := s.Store([]byte("Hello"))
	assert.NoError(t, err)
}
