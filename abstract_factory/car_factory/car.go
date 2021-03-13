package car_factory

import (
	"errors"
	"strings"
)

type (
	// body は車を維持します
	body interface {
		sustain() string
	}

	// engine は始動します
	engine interface {
		start() string
	}

	// tire は抵抗を提供します
	tire interface {
		grip() string
	}

	// Car は生成対象です
	Car struct {
		body
		engine
		tires []tire
	}
)

// Run は車を実際に走らせます
func (c Car) Run() (string, error) {
	if c.body == nil {
		return "", errors.New("error: no body")
	}
	if c.engine == nil {
		return "", errors.New("error: no engine")
	}
	if len(c.tires) != 4 {
		return "", errors.New("error: non 4 tires")
	}
	msgArray := []string{c.body.sustain(), c.engine.start()}
	for _, t := range c.tires {
		msgArray = append(msgArray, t.grip())
	}
	return strings.Join(msgArray, ","), nil
}
