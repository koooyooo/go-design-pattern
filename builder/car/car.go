package car

import "bytes"

type (
	car struct {
		engine Engine
		handle Handle
		tires  []Tire
	}
)

func (c car) Drive() string {
	var s bytes.Buffer
	s.WriteString(c.engine.Drive())
	s.WriteString(",")
	s.WriteString(c.handle.Control())
	for _, t := range c.tires {
		s.WriteString(",")
		s.WriteString(t.Grip())
	}
	return s.String()
}
