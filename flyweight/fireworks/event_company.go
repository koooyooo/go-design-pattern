package fireworks

import (
	"fmt"
	"time"
)

// eventCompany はイベント業者
type eventCompany struct {
	cnt int
}

// BorrowToilet は仮設トイレをレンタルする
func (f *eventCompany) BorrowToilet() *tmpToilet {
	t := &tmpToilet{ID: f.cnt}
	f.cnt++
	return t
}

// BorrowCnt は総レンタル数を取得する
func (f *eventCompany) BorrowCnt() int {
	return f.cnt
}

// tmpToilet は仮設トイレを表す構造体
type tmpToilet struct {
	ID int
}

// Use は仮設トイレを使用する（時間がかかる）
func (t *tmpToilet) Use() {
	time.Sleep(10 * time.Microsecond)
}

// String は仮設トイレの文字列表現
func (t *tmpToilet) String() string {
	return fmt.Sprintf("tmpToilet:(%d)", t.ID)
}
