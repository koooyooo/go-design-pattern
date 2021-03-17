package bridge

type (
	Human struct {
		Name string
		Job  Job
	}

	Job struct {
		Name  string
		Title string
	}
)
