// Generated SBE (Simple Binary Encoding) message codec

package deribit_multicast

import (
	"fmt"
	"io"
	"math"
	"time"
)

type Rfq struct {
	Header       MessageHeader
	InstrumentId uint32
	State        YesNoEnum
	Side         RfqDirectionEnum
	Amount       float64
	TimestampMs  uint64
}

func (r *Rfq) PPrint(i int) {
	PPrintlnInd(i, "Object: RFQ")
	r.Header.PPrint(i + 2)
	PPrintlnInd(i+2, "InstrumentId:", r.InstrumentId)
	PPrintlnInd(i+2, "State:", r.State.GetPPrint()) // TODO print nicer state message
	PPrintlnInd(i+2, "Side:", r.Side.GetPPrint())
	PPrintlnInd(i+2, "Amount:", r.Amount)
	PPrintlnInd(i+2, "TimestampMS:", time.UnixMilli(int64(r.TimestampMs)))
}

func (r *Rfq) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if err := _m.WriteUint32(_w, r.InstrumentId); err != nil {
		return err
	}
	if err := r.State.Encode(_m, _w); err != nil {
		return err
	}
	if err := r.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, r.Amount); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, r.TimestampMs); err != nil {
		return err
	}
	return nil
}

func (r *Rfq) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if err := _m.ReadUint32(_r, &r.InstrumentId); err != nil {
		return err
	}
	if err := r.State.Decode(_m, _r, actingVersion); err != nil {
		return err
	}
	if err := r.Side.Decode(_m, _r, actingVersion); err != nil {
		return err
	}
	if err := _m.ReadFloat64(_r, &r.Amount); err != nil {
		return err
	}
	if err := _m.ReadUint64(_r, &r.TimestampMs); err != nil {
		return err
	}
	return nil
}

func (r *Rfq) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if r.InstrumentIdInActingVersion(actingVersion) {
		if r.InstrumentId < r.InstrumentIdMinValue() || r.InstrumentId > r.InstrumentIdMaxValue() {
			return fmt.Errorf("range check failed on r.InstrumentId (%v < %v > %v)", r.InstrumentIdMinValue(), r.InstrumentId, r.InstrumentIdMaxValue())
		}
	}
	if err := r.State.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := r.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if r.AmountInActingVersion(actingVersion) {
		if r.Amount < r.AmountMinValue() || r.Amount > r.AmountMaxValue() {
			return fmt.Errorf("range check failed on r.Amount (%v < %v > %v)", r.AmountMinValue(), r.Amount, r.AmountMaxValue())
		}
	}
	if r.TimestampMsInActingVersion(actingVersion) {
		if r.TimestampMs < r.TimestampMsMinValue() || r.TimestampMs > r.TimestampMsMaxValue() {
			return fmt.Errorf("range check failed on r.TimestampMs (%v < %v > %v)", r.TimestampMsMinValue(), r.TimestampMs, r.TimestampMsMaxValue())
		}
	}
	return nil
}

func RfqInit(r *Rfq) {
}

func (*Rfq) SbeBlockLength() (blockLength uint16) {
	return 34
}

func (*Rfq) SbeTemplateId() (templateId uint16) {
	return 1009
}

func (*Rfq) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*Rfq) SbeSchemaVersion() (schemaVersion uint16) {
	return 3
}

func (*Rfq) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Rfq) SbeSemanticVersion() (semanticVersion string) {
	return ""
}

func (*Rfq) HeaderId() uint16 {
	return 1
}

func (*Rfq) HeaderSinceVersion() uint16 {
	return 0
}

func (r *Rfq) HeaderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.HeaderSinceVersion()
}

func (*Rfq) HeaderDeprecated() uint16 {
	return 0
}

func (*Rfq) HeaderMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Rfq) InstrumentIdId() uint16 {
	return 2
}

func (*Rfq) InstrumentIdSinceVersion() uint16 {
	return 0
}

func (r *Rfq) InstrumentIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.InstrumentIdSinceVersion()
}

func (*Rfq) InstrumentIdDeprecated() uint16 {
	return 0
}

func (*Rfq) InstrumentIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Rfq) InstrumentIdMinValue() uint32 {
	return 0
}

func (*Rfq) InstrumentIdMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*Rfq) InstrumentIdNullValue() uint32 {
	return math.MaxUint32
}

func (*Rfq) StateId() uint16 {
	return 3
}

func (*Rfq) StateSinceVersion() uint16 {
	return 0
}

func (r *Rfq) StateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.StateSinceVersion()
}

func (*Rfq) StateDeprecated() uint16 {
	return 0
}

func (*Rfq) StateMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Rfq) SideId() uint16 {
	return 4
}

func (*Rfq) SideSinceVersion() uint16 {
	return 0
}

func (r *Rfq) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SideSinceVersion()
}

func (*Rfq) SideDeprecated() uint16 {
	return 0
}

func (*Rfq) SideMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Rfq) AmountId() uint16 {
	return 5
}

func (*Rfq) AmountSinceVersion() uint16 {
	return 0
}

func (r *Rfq) AmountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.AmountSinceVersion()
}

func (*Rfq) AmountDeprecated() uint16 {
	return 0
}

func (*Rfq) AmountMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Rfq) AmountMinValue() float64 {
	return -math.MaxFloat64
}

func (*Rfq) AmountMaxValue() float64 {
	return math.MaxFloat64
}

func (*Rfq) AmountNullValue() float64 {
	return math.NaN()
}

func (*Rfq) TimestampMsId() uint16 {
	return 6
}

func (*Rfq) TimestampMsSinceVersion() uint16 {
	return 0
}

func (r *Rfq) TimestampMsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.TimestampMsSinceVersion()
}

func (*Rfq) TimestampMsDeprecated() uint16 {
	return 0
}

func (*Rfq) TimestampMsMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*Rfq) TimestampMsMinValue() uint64 {
	return 0
}

func (*Rfq) TimestampMsMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*Rfq) TimestampMsNullValue() uint64 {
	return math.MaxUint64
}
