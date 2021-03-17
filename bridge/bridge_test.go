// [目的]
// Bridgeはクラスを拡張させる際に 2つ以上の軸をもたせて制御させるパターンです。
//
// [Writer] + [Format]
// io.Stdin + JSON / Yaml
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

func TestBridgeStdoutJSON(t *testing.T) {
	// 標準出力に JSONとして出力
	stdoutJSON := bridge.WritingTarget{
		Write:  bridge.WriteStdout,
		Format: bridge.FormatJSON,
	}
	s, err := stdoutJSON.WriteOut(human)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, `[stdout] {"Name":"Bridge","Job":{"Name":"Bridge Architect","Title":"Manager"}}`, s)
}

func TestBridgeStdoutYAML(t *testing.T) {
	// 標準出力に YAMLとして出力
	stdoutYAML := bridge.WritingTarget{
		Write:  bridge.WriteStdout,
		Format: bridge.FormatYAML,
	}
	s, err := stdoutYAML.WriteOut(human)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, `[stdout] name: Bridge
job:
    name: Bridge Architect
    title: Manager
`, s)
}

func TestBridgeStderrJSON(t *testing.T) {
	// 標準エラー出力に JSONとして出力
	stderrJSON := bridge.WritingTarget{
		Write:  bridge.WriteStderr,
		Format: bridge.FormatJSON,
	}
	s, err := stderrJSON.WriteOut(human)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, `[stderr] {"Name":"Bridge","Job":{"Name":"Bridge Architect","Title":"Manager"}}`, s)
}

func TestBridgeStderrYAML(t *testing.T) {
	// 標準エラー出力に YAMLとして出力
	stderrYAML := bridge.WritingTarget{
		Write:  bridge.WriteStderr,
		Format: bridge.FormatYAML,
	}
	s, err := stderrYAML.WriteOut(human)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, `[stderr] name: Bridge
job:
    name: Bridge Architect
    title: Manager
`, s)
}
