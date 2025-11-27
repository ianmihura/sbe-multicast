package stdmsg

import (
	"time"
)

type Ticker struct {
	Header                 MessageHeader
	InstrumentId           uint32
	InstrumentState        InstrumentStateEnum
	TimestampMs            uint64
	OpenInterest           float64
	MinSellPrice           float64
	MaxBuyPrice            float64
	LastPrice              float64
	IndexPrice             float64
	MarkPrice              float64
	BestBidPrice           float64
	BestBidAmount          float64
	BestAskPrice           float64
	BestAskAmount          float64
	CurrentFunding         float64
	Funding8h              float64
	EstimatedDeliveryPrice float64
	DeliveryPrice          float64
	SettlementPrice        float64
}

func (m *Ticker) PPrint(i int) {
	PPrintlnInd(i, "Ticker")
	m.Header.PPrint(i + 2)
	PPrintlnInd(i+2, "InstrumentId:", m.InstrumentId)
	m.InstrumentState.PPrint(i + 2)
	PPrintlnInd(i+2, "TimestampMs:", time.UnixMilli(int64(m.TimestampMs)))
	PPrintlnInd(i+2, "OpenInterest:", m.OpenInterest)
	PPrintlnInd(i+2, "MinSellPrice:", m.MinSellPrice)
	PPrintlnInd(i+2, "MaxBuyPrice:", m.MaxBuyPrice)
	PPrintlnInd(i+2, "LastPrice:", m.LastPrice)
	PPrintlnInd(i+2, "IndexPrice:", m.IndexPrice)
	PPrintlnInd(i+2, "MarkPrice:", m.MarkPrice)
	PPrintlnInd(i+2, "BestBidPrice:", m.BestBidPrice)
	PPrintlnInd(i+2, "BestBidAmount:", m.BestBidAmount)
	PPrintlnInd(i+2, "BestAskPrice:", m.BestAskPrice)
	PPrintlnInd(i+2, "BestAskAmount:", m.BestAskAmount)
	PPrintlnInd(i+2, "CurrentFunding:", m.CurrentFunding)
	PPrintlnInd(i+2, "Funding8h:", m.Funding8h)
	PPrintlnInd(i+2, "EstimatedDeliveryPrice:", m.EstimatedDeliveryPrice)
	PPrintlnInd(i+2, "DeliveryPrice:", m.DeliveryPrice)
	PPrintlnInd(i+2, "SettlementPrice:", m.SettlementPrice)
}

func (m *Ticker) Decode(c *Coder) {
	c.Decode(&m.InstrumentId)
	c.Decode(&m.InstrumentState)
	c.Decode(&m.TimestampMs)
	c.Decode(&m.OpenInterest)
	c.Decode(&m.MinSellPrice)
	c.Decode(&m.MaxBuyPrice)
	c.Decode(&m.LastPrice)
	c.Decode(&m.IndexPrice)
	c.Decode(&m.MarkPrice)
	c.Decode(&m.BestBidPrice)
	c.Decode(&m.BestBidAmount)
	c.Decode(&m.BestAskPrice)
	c.Decode(&m.BestAskAmount)
	c.Decode(&m.CurrentFunding)
	c.Decode(&m.Funding8h)
	c.Decode(&m.EstimatedDeliveryPrice)
	c.Decode(&m.DeliveryPrice)
	c.Decode(&m.SettlementPrice)
}
