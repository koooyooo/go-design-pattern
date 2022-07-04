package compute

import "io"

type Initializer interface {
	init() error
}

type Finalizer interface {
	finalize() error
}

type CPU interface {
	Initializer
	Finalizer
	calculate([]byte) ([]byte, error)
}

type Memory interface {
	Initializer
	Finalizer
	io.ReadWriter
}

type HardDrive interface {
	Initializer
	Finalizer
	io.ReadWriter
}

type USBDevice interface {
	Initializer
	Finalizer
	io.ReadWriter
}

type HDMIDevice interface {
	Initializer
	Finalizer
	io.ReadWriter
}
