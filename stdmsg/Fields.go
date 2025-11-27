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

type GroupSizeEncoding struct {
	BlockLength      uint16
	NumInGroup       uint16
	NumGroups        uint16
	NumVarDataFields uint16
}

func (m *GroupSizeEncoding) PPrint(i int) {
	PPrintlnInd(i, "BlockLength:", m.BlockLength)
	PPrintlnInd(i, "NumInGroup:", m.NumInGroup)
	PPrintlnInd(i, "NumGroups:", m.NumGroups)
	PPrintlnInd(i, "NumVarDataFields:", m.NumVarDataFields)
}

func (m *GroupSizeEncoding) Decode(c *Coder) {
	c.Decode(&m.BlockLength)
	c.Decode(&m.NumInGroup)
	c.Decode(&m.NumGroups)
	c.Decode(&m.NumVarDataFields)
}
