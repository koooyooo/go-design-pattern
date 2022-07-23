// Package facade_pattern
//
// [目的]
// Facadeパターンの目的は、外部から見えるインターフェイスをシンプルにするものです。
//
// [概要]
// 例えばある処理を正確に実行するには、10個の関数を順序どおりに操作する必要があるとします。
// この様な複雑な操作を外部に提供すると、次の問題が発生します。
// - 利用側が正確に操作するために要求する知識の量が膨大になる
// - 利用側が正確に操作できない可能性が高くなる
//
// Facadeは複雑性を隠蔽するシンプルな窓口を提供し、利用側の利便性を高めると同時に安定的な利用フローを保証することができます。
//
package facade_pattern

import (
	"github.com/koooyooo/go-design-pattern/facade/vehicle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFacade(t *testing.T) {
	facade := vehicle.InspectionCompanyFacade{}
	errs := facade.CheckAll()
	assert.Equal(t, 0, len(errs))
}
