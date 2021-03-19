// [目的]
// Bridgeの目的はクラスを拡張時の拡張の方向性が2系統あった際でも、全組み合わせの実装を用意せずに最適化することです。
//
// [概要]
// 例えばクラスの役割が、ある出力先にあるフォーマットで出力することだとします。
// 可能なら、複数の出力先を選びたいし、複数のフォーマットも選びたい。しかしこれらを継承で用意するには全ての組み合わせ分のサブクラスが必要です。
// そこで、拡張方向の片系統は通常継承で実現しつつも、拡張方向のもう片系統は別クラスに任せます。これによって2方向の拡張が可能です。
// しかしこのままだと機能的に統合することができません。そこで、機能統合には継承ではなく委譲を使います。
// 2系統の流派を用意し、片方の流派がもう片方の流派に頼ることで、全ての機能を手に入れる、この継承と委譲のコンビネーションが Bridgeです。
//
// [実装]
// Goの場合は継承がありませんので、継承に当たる部分は関数の入れ替えで表現します。
// 委譲の部分も通常のクラスや抽象クラスをOverrideすることはできませんので、Interfaceによるポリモーフィズムで代替します。
// Javaのコンセプトに近づけてこの様な表現にしましたが、関数やInterfaceによる実装入れ替えは関数だけを用いても、Interfaceだけを用いても実現可能です。
//
// [シナリオ]
// 今回は 出力先と出力フォーマットの2種類の軸で拡張可能な Bridgeを用意し、出力先が出力フォーマットに委譲する形を取りました。
// 出力先は Stdoutと Stderrとし、テストの都合上出力内容は stringでも返すようにしました。
// 出力フォーマットは JSON, YAML を用意しました。
// 単なる継承の場合、2 x 2 の 4通りの継承クラスを用意する必要がありますが、Bridgeパターンを使えば実装の組み換えだけで同じことが可能です。
package bridge_pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/bridge/bridge"
)

// 出力対象の構造体を準備
var human = bridge.Human{
	Name: "Bridge",
	Job: bridge.Job{
		Name:  "Bridge Architect",
		Title: "Manager",
	},
}

type testcase struct {
	name     string
	model    bridge.Human
	target   *bridge.WritingTarget
	expected string
}

func (tc testcase) assert(t *testing.T) {
	s, err := tc.target.WriteOut(human)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, tc.expected, s)
}

func Test4Patterns(t *testing.T) {
	// 4系統の組み合わせを用意し、全ての組み合わせで稼働していることを確認
	stdoutJSON := &bridge.WritingTarget{
		Write:  bridge.WriteStdout,
		Format: bridge.FormatJSON,
	}
	stdoutYAML := &bridge.WritingTarget{
		Write:  bridge.WriteStdout,
		Format: bridge.FormatYAML,
	}
	stderrJSON := &bridge.WritingTarget{
		Write:  bridge.WriteStderr,
		Format: bridge.FormatJSON,
	}
	stderrYAML := &bridge.WritingTarget{
		Write:  bridge.WriteStderr,
		Format: bridge.FormatYAML,
	}

	cases := []testcase{
		{
			name:     "stdout x json",
			model:    human,
			target:   stdoutJSON,
			expected: `[stdout] {"Name":"Bridge","Job":{"Name":"Bridge Architect","Title":"Manager"}}`,
		},
		{
			name:   "stdout x yaml",
			model:  human,
			target: stdoutYAML,
			expected: `[stdout] name: Bridge
job:
    name: Bridge Architect
    title: Manager
`,
		},
		{
			name:     "stderr x json",
			model:    human,
			target:   stderrJSON,
			expected: `[stderr] {"Name":"Bridge","Job":{"Name":"Bridge Architect","Title":"Manager"}}`,
		},
		{
			name:   "stderr x yaml",
			model:  human,
			target: stderrYAML,
			expected: `[stderr] name: Bridge
job:
    name: Bridge Architect
    title: Manager
`,
		},
	}
	for _, c := range cases {
		c.assert(t)
	}
}
