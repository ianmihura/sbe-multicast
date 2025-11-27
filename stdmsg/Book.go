package stdmsg

import "time"

type Book struct {
	Header       MessageHeader
	InstrumentId uint32
	TimestampMs  uint64
	PrevChangeId uint64
	ChangeId     uint64
	IsLast       YesNoEnum
	ChangesList  GroupBookChangesList
}

func (m *Book) PPrint(i int) {
	PPrintlnInd(i, "Book")
	m.Header.PPrint(i + 2)
	PPrintlnInd(i+2, "InstrumentId:", m.InstrumentId)
	PPrintlnInd(i+2, "TimestampMs:", time.UnixMilli(int64(m.TimestampMs)))
	PPrintlnInd(i+2, "PrevChangeId:", m.PrevChangeId)
	PPrintlnInd(i+2, "ChangeId:", m.ChangeId)
	m.IsLast.PPrintCustom(i+2, "IsLast:", "more to come", "last")
	m.ChangesList.PPrint(i + 2)
}

func (m *Book) Decode(c *Coder) {
	c.Decode(&m.InstrumentId)
	c.Decode(&m.TimestampMs)
	c.Decode(&m.PrevChangeId)
	c.Decode(&m.ChangeId)
	c.Decode(&m.IsLast)
	m.ChangesList.Decode(c)
}

type GroupBookChangesList struct {
	GroupHeader GroupSizeEncoding
	Changes     []BookChangesItem
}

func (m *GroupBookChangesList) PPrint(i int) {
	PPrintlnInd(i, "BookChangesList")
	PPrintlnInd(i+2, "GroupHeader:")
	m.GroupHeader.PPrint(i + 4)
	PPrintlnInd(i+2, "ChangesList:")
	for i_ := range m.Changes {
		m.Changes[i_].PPrint(i + 4)
	}
}

func (m *GroupBookChangesList) Decode(c *Coder) {
	m.GroupHeader.Decode(c)
	m.Changes = make([]BookChangesItem, m.GroupHeader.NumInGroup)
	for i_ := range m.Changes {
		m.Changes[i_].Decode(c)
	}
}

type BookChangesItem struct {
	Side   BookSideEnum
	Change BookChangeEnum
	Price  float64
	Amount float64
}

func (m *BookChangesItem) PPrint(i int) {
	PPrintlnInd(i, "BookChangesItem")
	m.Side.PPrint(i + 2)
	m.Change.PPrint(i + 2)
	PPrintlnInd(i+2, "Price:", m.Price)
	PPrintlnInd(i+2, "Amount:", m.Amount)
}

func (m *BookChangesItem) Decode(c *Coder) {
	c.Decode(&m.Side)
	c.Decode(&m.Change)
	c.Decode(&m.Price)
	c.Decode(&m.Amount)
}
