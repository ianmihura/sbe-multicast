package stdmsg

import (
	"strconv"
)

type UnknownTemplateIdError uint16
type NotImplementedTemplateIdError uint16

func (e UnknownTemplateIdError) Error() string { return "unknown template id " + strconv.Itoa(int(e)) }
func (e NotImplementedTemplateIdError) Error() string {
	return "not implemented for template id " + strconv.Itoa(int(e))
}

type MessageHeader struct {
	BlockLength      uint16
	TemplateId       uint16
	SchemaId         uint16
	Version          uint16
	NumGroups        uint16
	NumVarDataFields uint16
	SequenceNumber   uint32 // from FrameHeader, not original from SBE MessageHeader
	Tmp              uint32 // Scratchtape variable for debugging
}

func (m *MessageHeader) Decode(c *Coder) {
	c.Decode(&m.BlockLength)
	c.Decode(&m.TemplateId)
	c.Decode(&m.SchemaId)
	c.Decode(&m.Version)
	c.Decode(&m.NumGroups)
	c.Decode(&m.NumVarDataFields)
}

// Returns a concrete StdMessage with the header attached to it
func (m *MessageHeader) GetConcreteMessage() (StdMessage, error) {
	switch m.TemplateId {
	case 1000:
		obj := &Instrument{}
		obj.Header = *m
		return obj, nil
	case 1001:
		obj := &Book{}
		obj.Header = *m
		return obj, nil
	case 1002:
		obj := &Trades{}
		obj.Header = *m
		return obj, nil
	case 1003:
		obj := &Ticker{}
		obj.Header = *m
		return obj, nil
	case 1004:
		obj := &Snapshot{}
		obj.Header = *m
		return obj, nil
	case 1005:
		obj := &SnapshotStart{}
		obj.Header = *m
		return obj, nil
	case 1006:
		obj := &SnapshotEnd{}
		obj.Header = *m
		return obj, nil
	case 1007:
		obj := &ComboLegs{}
		obj.Header = *m
		return obj, nil
	case 1008:
		obj := &PriceIndex{}
		obj.Header = *m
		return obj, nil
	case 1009:
		obj := &Rfq{}
		obj.Header = *m
		return obj, nil
	case 1010:
		obj := &InstrumentV2{}
		obj.Header = *m
		return obj, nil
	default:
		return nil, UnknownTemplateIdError(m.TemplateId)
	}
}

func (m *MessageHeader) PPrint(i int) {
	PPrintlnInd(i, "Message Header")
	PPrintlnInd(i+2, "BlockLength:", m.BlockLength)
	PPrintlnInd(i+2, "TemplateId:", m.getTemplateName())
	PPrintlnInd(i+2, "SchemaId:", m.SchemaId)
	PPrintlnInd(i+2, "Version:", m.Version)
	PPrintlnInd(i+2, "NumGroups:", m.NumGroups)
	PPrintlnInd(i+2, "NumVarDataFields:", m.NumVarDataFields)
	PPrintlnInd(i+2, "SequenceNumber:", m.SequenceNumber)
	PPrintlnInd(i+2, "-- Tmp:", m.Tmp)
}

func (m *MessageHeader) getTemplateName() string {
	switch m.TemplateId {
	case 1000:
		return "instrument (1000)"
	case 1001:
		return "book (1001)"
	case 1002:
		return "trades (1002)"
	case 1003:
		return "ticker (1003)"
	case 1004:
		return "snapshot (1004)"
	case 1005:
		return "snapshot_start (1005)"
	case 1006:
		return "snapshot_end (1006)"
	case 1007:
		return "combo_legs (1007)"
	case 1008:
		return "price_index (1008)"
	case 1009:
		return "rfq (1009)"
	case 1010:
		return "instrument_v2 (1010)"
	default:
		return "unknown"
	}
}
