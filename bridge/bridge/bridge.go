package bridge

type WritingBridge struct {
	WritingTarget
	WritingFormat
}

func (wb WritingBridge) WriteOut(v interface{}) error {
	b, err := wb.Marshal(v)
	if err != nil {
		return err
	}
	_, err = wb.Write(b)
	if err != nil {
		return err
	}
	return nil
}
