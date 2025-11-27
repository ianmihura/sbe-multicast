package stdmsg

import "time"

type Trades struct {
	Header       MessageHeader
	InstrumentId uint32
	TradesList   GroupTradesList
}

func (m *Trades) PPrint(i int) {
	PPrintlnInd(i, "Trades")
	m.Header.PPrint(i + 2)
	PPrintlnInd(i+2, "InstrumentId:", m.InstrumentId)
	m.TradesList.PPrint(i + 2)
}

func (m *Trades) Decode(c *Coder) {
	c.Decode(&m.InstrumentId)
	m.TradesList.Decode(c)
}

type GroupTradesList struct {
	GroupHeader GroupSizeEncoding
	TradeList   []TradeItem
}

func (m *GroupTradesList) PPrint(i int) {
	PPrintlnInd(i, "TradesList")
	PPrintlnInd(i+2, "GroupHeader:")
	m.GroupHeader.PPrint(i + 4)
	PPrintlnInd(i+2, "Trades:")
	for i_ := range m.TradeList {
		PPrintlnInd(i+4, "TradeItem:")
		m.TradeList[i_].PPrint(i + 4)
	}
}

func (m *GroupTradesList) Decode(c *Coder) {
	c.Decode(&m.GroupHeader)
	m.TradeList = make([]TradeItem, m.GroupHeader.NumInGroup)
	for i := 0; i < int(m.GroupHeader.NumInGroup); i++ {
		c.Decode(&m.TradeList[i])
	}
}

type TradeItem struct {
	Direction     DirectionEnum
	Price         float64
	Amount        float64
	TimestampMs   uint64
	MarkPrice     float64
	IndexPrice    float64
	TradeSeq      uint64
	TradeId       uint64
	TickDirection TickDirectionEnum
	Liquidation   LiquidationEnum
	Iv            float64
	BlockTradeId  uint64
	ComboTradeId  uint64
}

func (m *TradeItem) PPrint(i int) {
	m.Direction.PPrint(i + 2)
	PPrintlnInd(i+2, "Price:", m.Price)
	PPrintlnInd(i+2, "Amount:", m.Amount)
	PPrintlnInd(i+2, "TimestampMs:", time.UnixMilli(int64(m.TimestampMs)))
	PPrintlnInd(i+2, "MarkPrice:", m.MarkPrice)
	PPrintlnInd(i+2, "IndexPrice:", m.IndexPrice)
	PPrintlnInd(i+2, "TradeSeq:", m.TradeSeq)
	PPrintlnInd(i+2, "TradeId:", m.TradeId)
	m.TickDirection.PPrint(i + 2)
	m.Liquidation.PPrint(i + 2)
	PPrintlnInd(i+2, "Iv:", m.Iv)
	PPrintlnInd(i+2, "BlockTradeId:", m.BlockTradeId)
	PPrintlnInd(i+2, "ComboTradeId:", m.ComboTradeId)
}

func (m *TradeItem) Decode(c *Coder) {
	c.Decode(&m.Direction)
	c.Decode(&m.Price)
	c.Decode(&m.Amount)
	c.Decode(&m.TimestampMs)
	c.Decode(&m.MarkPrice)
	c.Decode(&m.IndexPrice)
	c.Decode(&m.TradeSeq)
	c.Decode(&m.TradeId)
	c.Decode(&m.TickDirection)
	c.Decode(&m.Liquidation)
	c.Decode(&m.Iv)
	c.Decode(&m.BlockTradeId)
	c.Decode(&m.ComboTradeId)
}

type TickDirectionEnum uint8
type TickDirectionValues struct {
	plus      TickDirectionEnum
	zeroplus  TickDirectionEnum
	minus     TickDirectionEnum
	zerominus TickDirectionEnum
	NullValue TickDirectionEnum
}

var TickDirection = TickDirectionValues{0, 1, 2, 3, 255}

func (m *TickDirectionEnum) PPrint(i int) {
	switch *m {
	case TickDirection.plus:
		PPrintlnInd(i, "TickDirection: plus")
	case TickDirection.zeroplus:
		PPrintlnInd(i, "TickDirection: zeroplus")
	case TickDirection.minus:
		PPrintlnInd(i, "TickDirection: minus")
	case TickDirection.zerominus:
		PPrintlnInd(i, "TickDirection: zerominus")
	default:
		PPrintlnInd(i, "TickDirection: null")
	}
}

type LiquidationEnum uint8
type LiquidationValues struct {
	none      LiquidationEnum
	maker     LiquidationEnum
	taker     LiquidationEnum
	both      LiquidationEnum
	NullValue LiquidationEnum
}

var Liquidation = LiquidationValues{0, 1, 2, 3, 255}

func (m *LiquidationEnum) PPrint(i int) {
	switch *m {
	case Liquidation.none:
		PPrintlnInd(i, "Liquidation: none")
	case Liquidation.maker:
		PPrintlnInd(i, "Liquidation: maker")
	case Liquidation.taker:
		PPrintlnInd(i, "Liquidation: taker")
	case Liquidation.both:
		PPrintlnInd(i, "Liquidation: both")
	default:
		PPrintlnInd(i, "Liquidation: null")
	}
}
