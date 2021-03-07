package car

type (
	Engine interface {
		Drive() string
		Stop() string
	}
	V8Engine struct{}
	V4Engine struct{}
)

func (e V8Engine) Drive() string {
	return "Drive(V8)"
}
func (e V8Engine) Stop() string {
	return "Stop(V8)"
}

func (e V4Engine) Drive() string {
	return "Drive(V4)"
}
func (e V4Engine) Stop() string {
	return "Stop(V4)"
}
