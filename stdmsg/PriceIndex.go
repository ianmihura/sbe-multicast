package stdmsg

import "time"

type PriceIndex struct {
	Header      MessageHeader
	IndexName   [16]byte
	Price       float64
	TimestampMs uint64
}

func (m *PriceIndex) PPrint(i int) {
	PPrintlnInd(i, "Price Index")
	m.Header.PPrint(i + 2)
	PPrintlnInd(i+2, "IndexName:", string(m.IndexName[:]))
	PPrintlnInd(i+2, "Prce:", m.Price)
	PPrintlnInd(i+2, "TimestampMs:", time.UnixMilli(int64(m.TimestampMs)))
}

func (m *PriceIndex) Decode(c *Coder) {
	c.Decode(&m.IndexName)
	c.Decode(&m.Price)
	c.Decode(&m.TimestampMs)
}
