// [目的]
// GuardedSuspension の目的はタスク消費の最適化です。
// タスクのキューに一工夫を施し、仕事が無くなれば即休憩、仕事が入ったならば即稼働という、最適稼働を実現します。
//
// [概要]
// GolangにはChannelが存在します。実はChannelは実に良く完成されたGuardedSuspensionです。
// ここではChannelを用いたGuardedSuspensionと、独自に作成したGuardedSuspensionの2通りで実現したいと思います。
package guarded_suspension_pattern

import "testing"

func TestGuardedSuspension(t *testing.T) {

}
