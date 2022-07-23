package vehicle

type testerInspection struct {
}

func (ti *testerInspection) checkBreaks() error {
	return nil
}

func (ti *testerInspection) checkExhaustGas() error {
	return nil
}

func (ti *testerInspection) checkSideSlip() error {
	return nil
}

func (ti *testerInspection) checkLights() error {
	return nil
}
