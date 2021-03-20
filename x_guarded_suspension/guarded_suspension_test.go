// [目的]
// GuardedSuspension の目的はタスク消費の最適化です。
// タスクのキューに同期処理を施し「仕事が無くなれば即休憩 -> 仕事が入れば即再開」という最適稼働を実現します。
//
// [概要]
// GuardedSuspension は主にQueueとして実装します。
// Queueに要素がない場合待機させ、要素が追加された際に要素解除の通知を出すことで上記の目的を果たします。
//
// この仕組みの多くはプログラミング言語に依存します。JavaやGolangはこの仕組が用意されています。
// - Java: java.lang.Object のメソッド - Object#wait(), Object#notify(), Object#notifyAll()
// - Golang: *sync.Cond のメソッド - c.Wait(), c.Signal(), c.Broadcast()
//
// また、これらのメソッドを実行する区域では同期で保護します
// - Java: synchronized 修飾子によるメソッド or ブロックの保護
// - Golang: *sync.Cond内に存在する *sync.Mutexによる m.Lock(), m.Unlock()
//
// また、上記とは別に 各言語には GuardedSuspensionの仕様を満たす完成品が存在します。
// - Java: java.util.Concurrent.ConcurrentLinkedQueue<E>
// - Golang: channel
//
// そのため、多くの場合は独自に用意する必要はなく、これらの既製品の活用で十分でしょう。
// ここでは、独自に作成したGuardedSuspensionを用いて基本的な実現方法を確認し、その後Channelを用いた実装でも動きを確認します。
//
// [実現]
// この実現には、独自のQueueを用意し、以下の実装を追加します。
// - Queueには RemoveFirst, AddLastの2メソッドを用意する
// - RemoveFirstは同期保護する
//   - RemoveFirstは保護区間内にガード条件を入れる、これはWait解除後にも繰り返し評価可能な様にループで実装する
//   - RemoveFirstはガード条件に合致した場合、ループ内で c.Wait()をコールし Threadを待機状態にする
// - AddFirstも同期保護する
//   - c.Signal(), c.Broadcast() 自体は同期不要だが 周辺の状態を守るために保護
//     - `It is allowed but not required for the caller to hold c.L during the call.` (c.Signal() & c.Broadcast())
//   - 要素追加後に待機中のThreadに対し再開を促す通知をする
//     - c.Signal() は待機中の 1Thread
//     - c.Broadcast() は待機中の 全Thread
package guarded_suspension_pattern

import (
	"sync"
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
	var wg sync.WaitGroup
	wg.Add(2)
	var logs []string
	// 10要素まで非同期で書き込めるようにバッファを用意
	q := make(chan interface{}, 10)

	// 要素追加を遅延起動(100ms後)
	go func() {
		time.Sleep(100 * time.Millisecond)
		q <- "first"
		logs = append(logs, "AddLast")
		wg.Done()
		time.Sleep(100 * time.Millisecond)
		q <- "second"
		logs = append(logs, "AddLast")
		wg.Done()
	}()

	// 要素が無い状態でコールするとその場で待機状態に入る
	// 先の遅延Threadから要素が追加されると処理が再開される
	first := <-q
	logs = append(logs, "RemoveFirst")
	second := <-q
	logs = append(logs, "RemoveFirst")

	wg.Wait()
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
	// 10要素まで非同期で書き込めるようにバッファを用意
	q := make(chan interface{}, 10)

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
