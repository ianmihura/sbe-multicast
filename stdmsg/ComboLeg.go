package stdmsg

type ComboLegs struct {
	Header       MessageHeader
	InstrumentId uint32
	LegsList     GroupComboLegsLegsList
}

type GroupComboLegsLegsList struct {
	GroupHeader GroupSizeEncoding
	Legs        []ComboLegsLegsItem
}

func (m *GroupComboLegsLegsList) PPrint(i int) {
	PPrintlnInd(i, "ComboLegsLegsList")
	PPrintlnInd(i+2, "GroupHeader:")
	m.GroupHeader.PPrint(i + 4)
	PPrintlnInd(i+2, "Legs:")
	for i_ := range m.Legs {
		m.Legs[i_].PPrint(i + 4)
	}
}

func (m *GroupComboLegsLegsList) Decode(c *Coder) {
	m.GroupHeader.Decode(c)
	m.Legs = make([]ComboLegsLegsItem, m.GroupHeader.NumInGroup)
	for i_ := range m.Legs {
		m.Legs[i_].Decode(c)
	}
}

type ComboLegsLegsItem struct {
	LegInstrumentId uint32
	LegSize         int32
}

func (m *ComboLegsLegsItem) PPrint(i int) {
	PPrintlnInd(i, "LegInstrumentId:", m.LegInstrumentId, "| LegSize:", m.LegSize)
}

func (m *ComboLegsLegsItem) Decode(c *Coder) {
	c.Decode(&m.LegInstrumentId)
	c.Decode(&m.LegSize)
}
