package car_factory

import "bytes"

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
	var runMsg bytes.Buffer
	runMsg.WriteString(c.body.sustain())
	runMsg.WriteString(",")
	runMsg.WriteString(c.engine.start())
	runMsg.WriteString(",")
	for i, t := range c.tires {
		if i != 0 {
			runMsg.WriteString(",")
		}
		runMsg.WriteString(t.grip())
	}
	return runMsg.String(), nil
}
