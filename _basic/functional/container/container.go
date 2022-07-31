package container

type Container[T any] interface {
	Unit(T) Container[T]
	Get() (*T, bool)
	Map2(a, b Container[any], f func(a, b any) any) Container[any] // [A,B,C] を指定したいが難しい
}
