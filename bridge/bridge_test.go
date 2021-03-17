// [目的]
// Bridgeはクラスを拡張させる際に 2つ以上の軸をもたせて制御させるパターンです。
//
// [Writer] + [Format]
// io.Stdin + JSON / Yaml
package bridge_pattern

import (
	"testing"

	"github.com/koooyooo/go-design-pattern/bridge/bridge"
)

func TestBridge(t *testing.T) {
	//
	human := bridge.Human{
		Name: "Bridge",
		Job: bridge.Job{
			Name:  "Bridge Architect",
			Title: "Manager",
		},
	}

	// 標準出力に JSONとして出力
	stdoutJSON := bridge.WritingBridge{
		bridge.TargetStdout,
		bridge.FormatJSON,
	}
	stdoutJSON.WriteOut(human)

	// 標準出力に YAMLとして出力
	stdoutYAML := bridge.WritingBridge{
		bridge.TargetStdout,
		bridge.FormatYAML,
	}
	stdoutYAML.WriteOut(human)

}
