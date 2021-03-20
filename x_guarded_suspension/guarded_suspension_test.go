// [目的]
// GuardedSuspension の目的はタスク消費の最適化です。
// タスクのキューに一工夫を施し、仕事が無くなれば即休憩、仕事が入ったならば即稼働という、最適稼働を実現します。
//
// [概要]
// GolangにはChannelが存在します。実はChannelは実に良く完成されたGuardedSuspensionです。
// ここではChannelを用いたGuardedSuspensionと、独自に作成したGuardedSuspensionの2通りで実現したいと思います。
package guarded_suspension_pattern

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/x_guarded_suspension/taskqueue"
)

func TestGuardedSuspensionWait(t *testing.T) {
	q := taskqueue.NewTaskQueue()
	// 要素追加を遅延起動(100ms後)
	go func() {
		time.Sleep(100 * time.Millisecond)
		q.AddLast("first")
		time.Sleep(100 * time.Millisecond)
		q.AddLast("second")
	}()
	// 要素が無い状態でコールするとその場で待機状態に入る (別の処理Threadから要素が追加されると処理が再開される)
	first := q.RemoveFirst()
	second := q.RemoveFirst()

	// ログを確認すると、待機-解除を繰り返しているのが確認できる
	// 1.最初に取得を試みるも待機状態となり、2.要素追加をすると、3.それを条件に待機状態を抜け要素取得に成功…を2回繰り返す
	assert.EqualValues(t, []string{
		"RemoveFirst-Wait",
		"AddLast",
		"RemoveFirst-Done",
		"RemoveFirst-Wait",
		"AddLast",
		"RemoveFirst-Done"}, q.Logs())
	assert.Equal(t, "first", first)
	assert.Equal(t, "second", second)
}

func TestGuardedSuspentionNoWait(t *testing.T) {
	q := taskqueue.NewTaskQueue()

	// 要素が既に存在する場合は、待機状態に入らず要素取得に成功する
	q.AddLast("first")
	q.AddLast("second")

	first := q.RemoveFirst()
	second := q.RemoveFirst()

	// ログを確認すると、待機なしに処理が完了しているのが確認できる
	assert.EqualValues(t, []string{
		"AddLast",
		"AddLast",
		"RemoveFirst-Done",
		"RemoveFirst-Done"}, q.Logs())
	assert.Equal(t, "first", first)
	assert.Equal(t, "second", second)
}
