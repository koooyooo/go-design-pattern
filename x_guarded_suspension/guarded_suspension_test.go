// [目的]
// GuardedSuspension の目的はタスク消費の最適化です。
// タスクのキューに同期処理を施し「仕事が無くなれば即休憩 -> 仕事が入れば即再開」という最適稼働を実現します。
//
// [概要]
//
// Golangには 既存のGuardedSuspensionが存在します。それがchannelです。
// ここでは独自に作成したGuardedSuspensionを用いて基本的な実現方法を確認し、その後Channelを用いて
package guarded_suspension_pattern

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/koooyooo/go-design-pattern/x_guarded_suspension/guarded_suspension"
)

// 独自実装を用いての待機が入るシナリオ
func TestGuardedSuspensionTaskQueueWait(t *testing.T) {
	q := guarded_suspension.NewTaskQueue()
	// 要素追加を遅延起動(100ms後)
	go func() {
		time.Sleep(100 * time.Millisecond)
		q.AddLast("first")
		time.Sleep(100 * time.Millisecond)
		q.AddLast("second")
	}()
	// 要素が無い状態でコールするとその場で待機状態に入る
	// 先の遅延Threadから要素が追加されると処理が再開される
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

// 独自実装を用いての待機が入らないシナリオ
func TestGuardedSuspensionTaskQueueNoWait(t *testing.T) {
	q := guarded_suspension.NewTaskQueue()

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

// Channel実装を用いての待機が入るシナリオ
func TestGuardedSuspensionChannelWait(t *testing.T) {
	var logs []string
	q := make(chan interface{})

	// 要素追加を遅延起動(100ms後)
	go func() {
		time.Sleep(100 * time.Millisecond)
		q <- "first"
		logs = append(logs, "AddLast")
		time.Sleep(100 * time.Millisecond)
		q <- "second"
		logs = append(logs, "AddLast")
	}()

	// 要素が無い状態でコールするとその場で待機状態に入る
	// 先の遅延Threadから要素が追加されると処理が再開される
	first := <-q
	logs = append(logs, "RemoveFirst")
	second := <-q
	logs = append(logs, "RemoveFirst")

	// ログを確認すると、処理順が異なるにも関わらず、追加-取得の順に繰り返しているのを確認できる (Waitはchannel内の制御のため出力できない)
	assert.EqualValues(t, []string{
		"AddLast",
		"RemoveFirst",
		"AddLast",
		"RemoveFirst"}, logs)
	assert.Equal(t, "first", first)
	assert.Equal(t, "second", second)
}

// Channel実装を用いての待機が入らないシナリオ
func TestGuardedSuspensionChannelNoWait(t *testing.T) {
	var logs []string
	q := make(chan interface{})

	q <- "first"
	logs = append(logs, "AddLast")
	q <- "second"
	logs = append(logs, "AddLast")

	// 要素が既に存在する場合は即時取得できる
	first := <-q
	logs = append(logs, "RemoveFirst")
	second := <-q
	logs = append(logs, "RemoveFirst")

	// ログを確認すると、待機が存在せず処理命令の順に実行されているのが確認できる
	assert.EqualValues(t, []string{
		"AddLast",
		"AddLast",
		"RemoveFirst",
		"RemoveFirst"}, logs)

	assert.Equal(t, "first", first)
	assert.Equal(t, "second", second)
}
