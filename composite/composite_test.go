// Package composite_pattern
//
// [目的]
// Compositeパターンの目的は複雑なツリー構造を意識せずシンプルに命令を飛ばすことです。
//
// [概要]
// 通常であれば、ツリー構造を意識して命令を飛ばさなければ、末端まで命令は届きません。
// しかし、ツリー上の構造物が次の役割を全うすれば、命令者は構造を全く意識せずに、末端まで命令を届けることができます。
//
// - 「枝は下部の枝葉に命令を伝えること」
// - 「葉は自身に命令を適用すること」
//
// Compositeパターンが活用可能な例は幾つかありますが、例えば以下のようなものです。
// 1. ファイルツリー全体に対しウイルススキャンする
// 2. 組織全体に対し命令を発する
// 3. ネットワーク全体の稼働をチェックする
//
// これらのケースのおける先述の担当は次の通りです。
// - 「枝は下部の枝葉に命令を伝えること」（1.ディレクトリ, 2.中間管理職, 3.ルーター/ハブ）
// - 「葉は自身に命令を適用すること」（1.ファイル, 2.メンバー, 3.ノード）
//
// [作成法]
// 1. 枝葉に共通する命令を持つインターフェイスを用意する
// 2. 枝葉で別々にこの命令を実現する実装を用意する
//   1. 枝ではループ処理で下位に命令を伝搬させる
//   2. 葉では命令を適用する
//
package composite_pattern

import (
	"github.com/koooyooo/go-design-pattern/composite/composite"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestComposite ではディレクトリ（testdir）に存在するテキストファイルの中身をすべて収集する過程をテストする。
//
// testdir
// ├── dir1
// │     └── dir2
// │         ├── bar.txt
// │         └── foo.txt
// ├── hello.txt
// └── world.txt
//
func TestComposite(t *testing.T) {
	// テストディレクトリに存在するファイルに対するファイルツリーを構築する
	tree, err := composite.NewTree("./testdir")
	assert.NoError(t, err)

	// ファイルツリーに対し構造を意識せずシンプルなオープン命令を出す
	contents, err := tree.Open()
	assert.NoError(t, err)

	// ツリー構造配下のテキストをすべて収集できているか確認する
	assert.Equal(t, []string{"Bar", "Foo", "Hello", "World"}, contents)
}
