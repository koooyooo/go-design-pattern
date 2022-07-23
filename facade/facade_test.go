// Package facade_pattern
//
// [目的]
// Facadeパターンの目的は、外部から見えるインターフェイスをシンプルにするものです。
//
// [概要]
// 例えばある処理を正確に実行するには、10個の関数を順序どおりに操作する必要があるとします。
// この様な複雑な操作を外部に提供するのは間違いの元です。
// Facadeはこれらを代行する窓口を提供します。
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
