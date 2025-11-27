package stdmsg

import "time"

type InstrumentV2 struct {
	Header                   MessageHeader
	InstrumentId             uint32
	InstrumentState          InstrumentStateEnum
	Kind                     InstrumentKindEnum
	InstrumentType           InstrumentTypeEnum
	OptionType               OptionTypeEnum
	SettlementPeriod         PeriodEnum
	SettlementPeriodCount    uint16
	BaseCurrency             [8]byte
	QuoteCurrency            [8]byte
	CounterCurrency          [8]byte
	SettlementCurrency       [8]byte
	SizeCurrency             [8]byte
	CreationTimestampMs      uint64
	ExpirationTimestampMs    uint64
	StrikePrice              float64
	ContractSize             float64
	MinTradeAmount           float64
	TickSize                 float64
	MakerCommission          float64
	TakerCommission          float64
	BlockTradeCommission     float64
	MaxLiquidationCommission float64
	MaxLeverage              float64
	TickStepsList            []InstrumentV2TickStepsList // TODO make
	InstrumentName           []uint8
}

type InstrumentV2TickStepsList struct {
	AbovePrice float64
	TickSize   float64
}

type Instrument struct {
	Header                   MessageHeader
	InstrumentId             uint32
	InstrumentState          InstrumentStateEnum
	Kind                     InstrumentKindEnum
	InstrumentType           InstrumentTypeEnum
	OptionType               OptionTypeEnum
	Rfq                      YesNoEnum
	SettlementPeriod         PeriodEnum
	SettlementPeriodCount    uint16
	BaseCurrency             [8]byte
	QuoteCurrency            [8]byte
	CounterCurrency          [8]byte
	SettlementCurrency       [8]byte
	SizeCurrency             [8]byte
	CreationTimestampMs      uint64
	ExpirationTimestampMs    uint64
	StrikePrice              float64
	ContractSize             float64
	MinTradeAmount           float64
	TickSize                 float64
	MakerCommission          float64
	TakerCommission          float64
	BlockTradeCommission     float64
	MaxLiquidationCommission float64
	MaxLeverage              float64
	InstrumentName           []uint8
}

func (m *InstrumentV2) PPrint(i int) {
	PPrintlnInd(i, "Instrument v2")
	m.Header.PPrint(i + 2)
	PPrintlnInd(i+2, "InstrumentId:", m.InstrumentId)
	m.InstrumentState.PPrint(i + 2)
	m.Kind.PPrint(i + 2)
	m.InstrumentType.PPrint(i + 2)
	m.OptionType.PPrint(i + 2)
	m.SettlementPeriod.PPrint(i + 2)
	PPrintlnInd(i+2, "SettlementPeriodCount:", m.SettlementPeriodCount)
	PPrintlnInd(i+2, "BaseCurrency:", string(m.BaseCurrency[:]))
	PPrintlnInd(i+2, "QuoteCurrency:", string(m.QuoteCurrency[:]))
	PPrintlnInd(i+2, "CounterCurrency:", string(m.CounterCurrency[:]))
	PPrintlnInd(i+2, "SettlementCurrency:", string(m.SettlementCurrency[:]))
	PPrintlnInd(i+2, "SizeCurrency:", string(m.SizeCurrency[:]))
	PPrintlnInd(i+2, "CreationTimestampMs:", time.UnixMilli(int64(m.CreationTimestampMs)))
	PPrintlnInd(i+2, "ExpirationTimestampMs:", time.UnixMilli(int64(m.ExpirationTimestampMs)))
	PPrintlnInd(i+2, "StrikePrice:", m.StrikePrice)
	PPrintlnInd(i+2, "ContractSize:", m.ContractSize)
	PPrintlnInd(i+2, "MinTradeAmount:", m.MinTradeAmount)
	PPrintlnInd(i+2, "TickSize:", m.TickSize)
	PPrintlnInd(i+2, "MakerCommission:", m.MakerCommission)
	PPrintlnInd(i+2, "TakerCommission:", m.TakerCommission)
	PPrintlnInd(i+2, "BlockTradeCommission:", m.BlockTradeCommission)
	PPrintlnInd(i+2, "MaxLiquidationCommission:", m.MaxLiquidationCommission)
	PPrintlnInd(i+2, "MaxLeverage:", m.MaxLeverage)
	// TODO TickStepsList
	PPrintlnInd(i+2, "InstrumentName:", m.InstrumentName)
}

func (m *Instrument) PPrint(i int) {
	PPrintlnInd(i, "Instrument")
	m.Header.PPrint(i + 2)
	PPrintlnInd(i+2, "InstrumentId:", m.InstrumentId)
	m.InstrumentState.PPrint(i + 2)
	m.Kind.PPrint(i + 2)
	m.InstrumentType.PPrint(i + 2)
	m.OptionType.PPrint(i + 2)
	m.Rfq.PPrint(i + 2)
	m.SettlementPeriod.PPrint(i + 2)
	PPrintlnInd(i+2, "SettlementPeriodCount:", m.SettlementPeriodCount)
	PPrintlnInd(i+2, "BaseCurrency:", string(m.BaseCurrency[:]))
	PPrintlnInd(i+2, "QuoteCurrency:", string(m.QuoteCurrency[:]))
	PPrintlnInd(i+2, "CounterCurrency:", string(m.CounterCurrency[:]))
	PPrintlnInd(i+2, "SettlementCurrency:", string(m.SettlementCurrency[:]))
	PPrintlnInd(i+2, "SizeCurrency:", string(m.SizeCurrency[:]))
	PPrintlnInd(i+2, "CreationTimestampMs:", time.UnixMilli(int64(m.CreationTimestampMs)))
	PPrintlnInd(i+2, "ExpirationTimestampMs:", time.UnixMilli(int64(m.ExpirationTimestampMs)))
	PPrintlnInd(i+2, "StrikePrice:", m.StrikePrice)
	PPrintlnInd(i+2, "ContractSize:", m.ContractSize)
	PPrintlnInd(i+2, "MinTradeAmount:", m.MinTradeAmount)
	PPrintlnInd(i+2, "TickSize:", m.TickSize)
	PPrintlnInd(i+2, "MakerCommission:", m.MakerCommission)
	PPrintlnInd(i+2, "TakerCommission:", m.TakerCommission)
	PPrintlnInd(i+2, "BlockTradeCommission:", m.BlockTradeCommission)
	PPrintlnInd(i+2, "MaxLiquidationCommission:", m.MaxLiquidationCommission)
	PPrintlnInd(i+2, "MaxLeverage:", m.MaxLeverage)
	PPrintlnInd(i+2, "InstrumentName:", m.InstrumentName)
}

func (m *InstrumentV2) Decode(c *Coder) {
	c.Decode(&m.InstrumentId)
	c.Decode(&m.InstrumentState)
	c.Decode(&m.Kind)
	c.Decode(&m.InstrumentType)
	c.Decode(&m.OptionType)
	c.Decode(&m.SettlementPeriod)
	c.Decode(&m.SettlementPeriodCount)
	c.Decode(&m.BaseCurrency)
	c.Decode(&m.QuoteCurrency)
	c.Decode(&m.CounterCurrency)
	c.Decode(&m.SettlementCurrency)
	c.Decode(&m.SizeCurrency)
	c.Decode(&m.CreationTimestampMs)
	c.Decode(&m.ExpirationTimestampMs)
	c.Decode(&m.StrikePrice)
	c.Decode(&m.ContractSize)
	c.Decode(&m.MinTradeAmount)
	c.Decode(&m.TickSize)
	c.Decode(&m.MakerCommission)
	c.Decode(&m.TakerCommission)
	c.Decode(&m.BlockTradeCommission)
	c.Decode(&m.MaxLiquidationCommission)
	c.Decode(&m.MaxLeverage)
	// TODO TickStepsList
	c.Decode(&m.InstrumentName) // TODO decode var length
}

func (m *Instrument) Decode(c *Coder) {
	c.Decode(&m.InstrumentId)
	c.Decode(&m.InstrumentState)
	c.Decode(&m.Kind)
	c.Decode(&m.InstrumentType)
	c.Decode(&m.OptionType)
	c.Decode(&m.Rfq)
	c.Decode(&m.SettlementPeriod)
	c.Decode(&m.SettlementPeriodCount)
	c.Decode(&m.BaseCurrency)
	c.Decode(&m.QuoteCurrency)
	c.Decode(&m.CounterCurrency)
	c.Decode(&m.SettlementCurrency)
	c.Decode(&m.SizeCurrency)
	c.Decode(&m.CreationTimestampMs)
	c.Decode(&m.ExpirationTimestampMs)
	c.Decode(&m.StrikePrice)
	c.Decode(&m.ContractSize)
	c.Decode(&m.MinTradeAmount)
	c.Decode(&m.TickSize)
	c.Decode(&m.MakerCommission)
	c.Decode(&m.TakerCommission)
	c.Decode(&m.BlockTradeCommission)
	c.Decode(&m.MaxLiquidationCommission)
	c.Decode(&m.MaxLeverage)
	c.Decode(&m.InstrumentName) // TODO decode var length
}

type InstrumentStateEnum uint8
type InstrumentStateValues struct {
	created     InstrumentStateEnum
	open        InstrumentStateEnum
	closed      InstrumentStateEnum
	settled     InstrumentStateEnum
	deactivated InstrumentStateEnum
	inactive    InstrumentStateEnum
	started     InstrumentStateEnum
	NullValue   InstrumentStateEnum
}

var InstrumentState = InstrumentStateValues{0, 1, 2, 3, 4, 5, 6, 255}

func (m *InstrumentStateEnum) PPrint(i int) {
	switch *m {
	case InstrumentState.created:
		PPrintlnInd(i, "State: created")
	case InstrumentState.open:
		PPrintlnInd(i, "State: open")
	case InstrumentState.closed:
		PPrintlnInd(i, "State: closed")
	case InstrumentState.settled:
		PPrintlnInd(i, "State: settled")
	case InstrumentState.deactivated:
		PPrintlnInd(i, "State: deactivated")
	case InstrumentState.inactive:
		PPrintlnInd(i, "State: inactive")
	case InstrumentState.started:
		PPrintlnInd(i, "State: started")
	default:
		PPrintlnInd(i, "State: null")
	}
}

type InstrumentKindEnum uint8
type InstrumentKindValues struct {
	future       InstrumentKindEnum
	option       InstrumentKindEnum
	future_combo InstrumentKindEnum
	option_combo InstrumentKindEnum
	spot         InstrumentKindEnum
	NullValue    InstrumentKindEnum
}

var InstrumentKind = InstrumentKindValues{0, 1, 2, 3, 4, 255}

func (m *InstrumentKindEnum) PPrint(i int) {
	switch *m {
	case InstrumentKind.future:
		PPrintlnInd(i, "Side: future")
	case InstrumentKind.option:
		PPrintlnInd(i, "Side: option")
	case InstrumentKind.future_combo:
		PPrintlnInd(i, "Side: future_combo")
	case InstrumentKind.option_combo:
		PPrintlnInd(i, "Side: option_combo")
	case InstrumentKind.spot:
		PPrintlnInd(i, "Side: spot")
	default:
		PPrintlnInd(i, "Side: null")
	}
}

type InstrumentTypeEnum uint8
type InstrumentTypeValues struct {
	not_applicable InstrumentTypeEnum
	reversed       InstrumentTypeEnum
	linear         InstrumentTypeEnum
	NullValue      InstrumentTypeEnum
}

var InstrumentType = InstrumentTypeValues{0, 1, 2, 255}

func (m *InstrumentTypeEnum) PPrint(i int) {
	switch *m {
	case InstrumentType.not_applicable:
		PPrintlnInd(i, "InstrumentType: not_applicable")
	case InstrumentType.reversed:
		PPrintlnInd(i, "InstrumentType: reversed")
	case InstrumentType.linear:
		PPrintlnInd(i, "InstrumentType: linear")
	default:
		PPrintlnInd(i, "InstrumentType: null")
	}
}

type OptionTypeEnum uint8
type OptionTypeValues struct {
	not_applicable OptionTypeEnum
	call           OptionTypeEnum
	put            OptionTypeEnum
	NullValue      OptionTypeEnum
}

var OptionType = OptionTypeValues{0, 1, 2, 255}

func (m *OptionTypeEnum) PPrint(i int) {
	switch *m {
	case OptionType.not_applicable:
		PPrintlnInd(i, "OptionType: not_applicable")
	case OptionType.call:
		PPrintlnInd(i, "OptionType: call")
	case OptionType.put:
		PPrintlnInd(i, "OptionType: put")
	default:
		PPrintlnInd(i, "OptionType: null")
	}
}

type PeriodEnum uint8
type PeriodValues struct {
	perpetual PeriodEnum
	minute    PeriodEnum
	hour      PeriodEnum
	day       PeriodEnum
	week      PeriodEnum
	month     PeriodEnum
	year      PeriodEnum
	NullValue PeriodEnum
}

var Period = PeriodValues{0, 1, 2, 3, 4, 5, 6, 255}

func (m *PeriodEnum) PPrint(i int) {
	switch *m {
	case Period.perpetual:
		PPrintlnInd(i, "Period: perpetual")
	case Period.minute:
		PPrintlnInd(i, "Period: minute")
	case Period.hour:
		PPrintlnInd(i, "Period: hour")
	case Period.day:
		PPrintlnInd(i, "Period: day")
	case Period.week:
		PPrintlnInd(i, "Period: week")
	case Period.month:
		PPrintlnInd(i, "Period: month")
	case Period.year:
		PPrintlnInd(i, "Period: year")
	default:
		PPrintlnInd(i, "Period: null")
	}
}
