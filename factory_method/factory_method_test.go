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
package factory_method

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/factory_method/factory"
)

// FactoryMethodを適用した場合
func TestFactoryMethod(t *testing.T) {
	// Factory経由で Storageを取得するがその実体は知らない（疎結合）
	s := factory.GetStorage()
	err := s.Store([]byte("Hello"))
	assert.NoError(t, err)
}

// FactoryMethodを適用しない場合
func TestNonFactory(t *testing.T) {
	var s factory.Storage
	// NopeStorageが実体だとバレてしまう (密結合)
	s = &factory.NopeStorage{}
	err := s.Store([]byte("Hello"))
	assert.NoError(t, err)
}
