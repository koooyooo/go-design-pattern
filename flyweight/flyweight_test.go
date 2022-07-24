// Package flyweight_pattern
//
// [目的]
// Flyweightパターンの目的はインスタンスをプーリングして再利用することで、インスタンスの生成コストを減らすことです。
// 特に生成に時間やリソースが必要なインスタンスの生成に対し適用される傾向があります。
//
package flyweight_pattern

import (
	"fmt"
	"github.com/koooyooo/go-design-pattern/flyweight/fireworks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFlyweight(t *testing.T) {
	// 仮設トイレゾーンを用意し仮設トイレを100個準備（足らなければ随時調達される）
	zone := fireworks.NewEventToiletZone()
	zone.ReadyToilet(100)

	// 10万回、順番に仮設トイレが利用される
	for i := 0; i < 100_000; i++ {
		time.Sleep(5 * time.Microsecond)
		go func() {
			zone.AssignToilet()
		}()
	}
	// 仮設トイレは適宜再利用されるため、最終的な調達回数は10万回以下となる（実際には100程度）
	fmt.Printf("Borrow Count: %d\n", zone.Company.BorrowCnt())
	assert.True(t, zone.Company.BorrowCnt() < 100_000)
}
