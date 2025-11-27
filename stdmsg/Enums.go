package stdmsg

type YesNoEnum uint8
type YesNoValues struct {
	No        YesNoEnum
	Yes       YesNoEnum
	NullValue YesNoEnum
}

var YesNo = YesNoValues{0, 1, 255}

func (m *YesNoEnum) PPrint(i int) {
	switch *m {
	case YesNo.No:
		PPrintlnInd(i, "no")
	case YesNo.Yes:
		PPrintlnInd(i, "yes")
	default:
		PPrintlnInd(i, "null")
	}
}

func (m *YesNoEnum) PPrintCustom(i int, prep, noStr, yesStr string) {
	switch *m {
	case YesNo.No:
		PPrintlnInd(i, prep, noStr)
	case YesNo.Yes:
		PPrintlnInd(i, prep, yesStr)
	default:
		PPrintlnInd(i, prep, "null")
	}
}

func (m *YesNoEnum) Decode(c *Coder) {
	c.Decode(m)
}
