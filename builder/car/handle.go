package car

type (
	Handle interface {
		Control() string
	}
	SportHandle  struct{}
	NormalHandle struct{}
)

func (h SportHandle) Control() string {
	return "Control(Sport)"
}

func (h NormalHandle) Control() string {
	return "Control(Normal)"
}
