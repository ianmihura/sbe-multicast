package stdmsg

import "time"

type Snapshot struct {
	Header         MessageHeader
	InstrumentId   uint32
	TimestampMs    uint64
	ChangeId       uint64
	IsBookComplete YesNoEnum
	IsLastInBook   YesNoEnum
	LevelsList     GroupSnapshotLevelsList
}

func (m *Snapshot) PPrint(i int) {
	PPrintlnInd(i, "Snapshot")
	m.Header.PPrint(i + 2)
	PPrintlnInd(i+2, "InstrumentId:", m.InstrumentId)
	PPrintlnInd(i+2, "TimestampMs:", time.UnixMilli(int64(m.TimestampMs)))
	PPrintlnInd(i+2, "ChangeId:", m.ChangeId)
	m.IsBookComplete.PPrintCustom(i+2, "IsBookComplete:", "incomplete (depth limited)", "complete")
	m.IsLastInBook.PPrintCustom(i+2, "IsLastInBook:", "more to come", "last")
	m.LevelsList.PPrint(i + 2)
}

func (m *Snapshot) Decode(c *Coder) {
	c.Decode(&m.InstrumentId)
	c.Decode(&m.TimestampMs)
	c.Decode(&m.ChangeId)
	c.Decode(&m.IsBookComplete)
	c.Decode(&m.IsLastInBook)
	m.LevelsList.Decode(c)
}

type GroupSnapshotLevelsList struct {
	GroupHeader GroupSizeEncoding
	Levels      []SnapshotLevelsItem
}

func (m *GroupSnapshotLevelsList) PPrint(i int) {
	PPrintlnInd(i, "SnapshotLevelsList")
	PPrintlnInd(i+2, "GroupHeader:")
	m.GroupHeader.PPrint(i + 4)
	PPrintlnInd(i+2, "LevelsList:")
	for i_ := range m.Levels {
		m.Levels[i_].PPrint(i + 4)
	}
}

func (m *GroupSnapshotLevelsList) Decode(c *Coder) {
	m.GroupHeader.Decode(c)
	m.Levels = make([]SnapshotLevelsItem, m.GroupHeader.NumInGroup)
	for i_ := range m.Levels {
		m.Levels[i_].Decode(c)
	}
}

type SnapshotLevelsItem struct {
	Side   BookSideEnum
	Price  float64
	Amount float64
}

func (m *SnapshotLevelsItem) PPrint(i int) {
	PPrintlnInd(i, "Side:", m.Side.GetPPrint(), "| Price:", m.Price, "| Amount:", m.Amount)
}

func (m *SnapshotLevelsItem) Decode(c *Coder) {
	c.Decode(&m.Side)
	c.Decode(&m.Price)
	c.Decode(&m.Amount)
}

type SnapshotStart struct {
	Header        MessageHeader
	SnapshotDelay uint32
}

func (m *SnapshotStart) PPrint(i int) {
	PPrintlnInd(i, "SnapshotStart")
	m.Header.PPrint(i + 2)
	PPrintlnInd(i+2, "SnapshotDelay:", m.SnapshotDelay)
}

func (m *SnapshotStart) Decode(c *Coder) {
	c.Decode(&m.SnapshotDelay)
}

type SnapshotEnd struct {
	Header MessageHeader
}

func (m *SnapshotEnd) PPrint(i int) {
	PPrintlnInd(i, "SnapshotEnd")
	m.Header.PPrint(i + 2)
}

func (m *SnapshotEnd) Decode(c *Coder) {
}
