package vehicle

// InspectionCompanyFacade は車検業者によるファサード
type InspectionCompanyFacade struct{}

//
func (icf *InspectionCompanyFacade) CheckAll() []error {
	var errors []error
	// 外観と内装の検査
	vi := &outlookInspection{}
	if err := vi.checkIdentity(); err != nil {
		errors = append(errors, err)
	}
	if err := vi.checkTires(); err != nil {
		errors = append(errors, err)
	}
	if err := vi.checkWindowShields(); err != nil {
		errors = append(errors, err)
	}
	if err := vi.checkMeters(); err != nil {
		errors = append(errors, err)
	}
	if err := vi.checkInteriorParts(); err != nil {
		errors = append(errors, err)
	}

	// 外回りの検査
	oi := &outsideInspection{}
	if err := oi.checkLights(); err != nil {
		errors = append(errors, err)
	}
	if err := oi.checkWipers(); err != nil {
		errors = append(errors, err)
	}
	if err := oi.checkWindowWasherFluid(); err != nil {
		errors = append(errors, err)
	}
	if err := oi.checkMuffler(); err != nil {
		errors = append(errors, err)
	}
	if err := oi.checkDriveShaft(); err != nil {
		errors = append(errors, err)
	}
	if err := oi.checkSteeringRackBoots(); err != nil {
		errors = append(errors, err)
	}

	// テスターによる検査
	ti := &testerInspection{}
	if err := ti.checkBreaks(); err != nil {
		errors = append(errors, err)
	}
	if err := ti.checkExhaustGas(); err != nil {
		errors = append(errors, err)
	}
	if err := ti.checkSideSlip(); err != nil {
		errors = append(errors, err)
	}
	if err := ti.checkLights(); err != nil {
		errors = append(errors, err)
	}
	return errors
}
