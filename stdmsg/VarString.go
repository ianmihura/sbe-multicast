package stdmsg

type VarString struct {
	Length  uint8
	VarData []uint8
}

func (m *VarString) PPrint(i int) {
	PPrintlnInd(i, "Length:", m.Length)
	PPrintlnInd(i, "VarData:", string(m.VarData[:]))
}

func (m *VarString) PPrintCustom(i int, prep string) {
	PPrintlnInd(i, prep, string(m.VarData[:]))
}

func (m *VarString) Decode(c *Coder) {
	c.Decode(&m.Length)
	m.VarData = make([]uint8, m.Length)
	c.Decode(&m.VarData)
}
