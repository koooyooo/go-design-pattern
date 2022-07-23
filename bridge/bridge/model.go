package bridge

type (
	// Human は人間を表す構造体
	Human struct {
		Name string
		Job  Job
	}
	// Job は職業を表す構造体
	Job struct {
		Name  string
		Title string
	}
)
