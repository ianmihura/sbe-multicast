// Generated SBE (Simple Binary Encoding) message codec

package deribit_multicast

import (
	"fmt"
	"io"
	"math"
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
}

func (h *MessageHeader) PPrint(i int) {
	PPrintlnInd(i, "Message Header")
	PPrintlnInd(i+2, "BlockLength:", h.BlockLength)
	PPrintlnInd(i+2, "TemplateId:", h.getTemplateName())
	PPrintlnInd(i+2, "SchemaId:", h.SchemaId)
	PPrintlnInd(i+2, "Version:", h.Version)
	PPrintlnInd(i+2, "NumGroups:", h.NumGroups)
	PPrintlnInd(i+2, "NumVarDataFields:", h.NumVarDataFields)
}

func (h *MessageHeader) GetConcreteMessage() (SbeStdMessage, error) {
	switch h.TemplateId {
	case 1000:
		return nil, NotImplementedTemplateIdError(h.TemplateId)
	case 1001:
		return nil, NotImplementedTemplateIdError(h.TemplateId)
	case 1002:
		return nil, NotImplementedTemplateIdError(h.TemplateId)
	case 1003:
		return nil, NotImplementedTemplateIdError(h.TemplateId)
	case 1004:
		return nil, NotImplementedTemplateIdError(h.TemplateId)
	case 1005:
		return nil, NotImplementedTemplateIdError(h.TemplateId)
	case 1006:
		return nil, NotImplementedTemplateIdError(h.TemplateId)
	case 1007:
		return nil, NotImplementedTemplateIdError(h.TemplateId)
	case 1008:
		return nil, NotImplementedTemplateIdError(h.TemplateId)
	case 1009:
		obj := &Rfq{}
		obj.Header = *h
		return obj, nil
	case 1010:
		return nil, NotImplementedTemplateIdError(h.TemplateId)
	default:
		return nil, UnknownTemplateIdError(h.TemplateId)
	}
}

func (h *MessageHeader) getTemplateName() string {
	switch h.TemplateId {
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

func (m *MessageHeader) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, m.BlockLength); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.TemplateId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.SchemaId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.Version); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.NumGroups); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.NumVarDataFields); err != nil {
		return err
	}
	return nil
}

func (m *MessageHeader) Decode(_m *SbeGoMarshaller, _r io.Reader) error {
	if _, err := io.ReadFull(_r, _m.b); err != nil {
		return err
	}
	m.BlockLength = uint16(_m.b[0]) | uint16(_m.b[1])<<8
	m.TemplateId = uint16(_m.b[2]) | uint16(_m.b[3])<<8
	m.SchemaId = uint16(_m.b[4]) | uint16(_m.b[5])<<8
	m.Version = uint16(_m.b[6]) | uint16(_m.b[7])<<8
	m.NumGroups = uint16(_m.b[8]) | uint16(_m.b[9])<<8
	m.NumVarDataFields = uint16(_m.b[10]) | uint16(_m.b[11])<<8
	return nil
}

func (m *MessageHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.BlockLengthInActingVersion(actingVersion) {
		if m.BlockLength < m.BlockLengthMinValue() || m.BlockLength > m.BlockLengthMaxValue() {
			return fmt.Errorf("range check failed on m.BlockLength (%v < %v > %v)", m.BlockLengthMinValue(), m.BlockLength, m.BlockLengthMaxValue())
		}
	}
	if m.TemplateIdInActingVersion(actingVersion) {
		if m.TemplateId < m.TemplateIdMinValue() || m.TemplateId > m.TemplateIdMaxValue() {
			return fmt.Errorf("range check failed on m.TemplateId (%v < %v > %v)", m.TemplateIdMinValue(), m.TemplateId, m.TemplateIdMaxValue())
		}
	}
	if m.SchemaIdInActingVersion(actingVersion) {
		if m.SchemaId < m.SchemaIdMinValue() || m.SchemaId > m.SchemaIdMaxValue() {
			return fmt.Errorf("range check failed on m.SchemaId (%v < %v > %v)", m.SchemaIdMinValue(), m.SchemaId, m.SchemaIdMaxValue())
		}
	}
	if m.VersionInActingVersion(actingVersion) {
		if m.Version < m.VersionMinValue() || m.Version > m.VersionMaxValue() {
			return fmt.Errorf("range check failed on m.Version (%v < %v > %v)", m.VersionMinValue(), m.Version, m.VersionMaxValue())
		}
	}
	if m.NumGroupsInActingVersion(actingVersion) {
		if m.NumGroups < m.NumGroupsMinValue() || m.NumGroups > m.NumGroupsMaxValue() {
			return fmt.Errorf("range check failed on m.NumGroups (%v < %v > %v)", m.NumGroupsMinValue(), m.NumGroups, m.NumGroupsMaxValue())
		}
	}
	if m.NumVarDataFieldsInActingVersion(actingVersion) {
		if m.NumVarDataFields < m.NumVarDataFieldsMinValue() || m.NumVarDataFields > m.NumVarDataFieldsMaxValue() {
			return fmt.Errorf("range check failed on m.NumVarDataFields (%v < %v > %v)", m.NumVarDataFieldsMinValue(), m.NumVarDataFields, m.NumVarDataFieldsMaxValue())
		}
	}
	return nil
}

func MessageHeaderInit(m *MessageHeader) {
}

func (*MessageHeader) EncodedLength() int64 {
	return 12
}

func (*MessageHeader) BlockLengthMinValue() uint16 {
	return 0
}

func (*MessageHeader) BlockLengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MessageHeader) BlockLengthNullValue() uint16 {
	return math.MaxUint16
}

func (*MessageHeader) BlockLengthSinceVersion() uint16 {
	return 0
}

func (m *MessageHeader) BlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BlockLengthSinceVersion()
}

func (*MessageHeader) BlockLengthDeprecated() uint16 {
	return 0
}

func (*MessageHeader) TemplateIdMinValue() uint16 {
	return 0
}

func (*MessageHeader) TemplateIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MessageHeader) TemplateIdNullValue() uint16 {
	return math.MaxUint16
}

func (*MessageHeader) TemplateIdSinceVersion() uint16 {
	return 0
}

func (m *MessageHeader) TemplateIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TemplateIdSinceVersion()
}

func (*MessageHeader) TemplateIdDeprecated() uint16 {
	return 0
}

func (*MessageHeader) SchemaIdMinValue() uint16 {
	return 0
}

func (*MessageHeader) SchemaIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MessageHeader) SchemaIdNullValue() uint16 {
	return math.MaxUint16
}

func (*MessageHeader) SchemaIdSinceVersion() uint16 {
	return 0
}

func (m *MessageHeader) SchemaIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SchemaIdSinceVersion()
}

func (*MessageHeader) SchemaIdDeprecated() uint16 {
	return 0
}

func (*MessageHeader) VersionMinValue() uint16 {
	return 0
}

func (*MessageHeader) VersionMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MessageHeader) VersionNullValue() uint16 {
	return math.MaxUint16
}

func (*MessageHeader) VersionSinceVersion() uint16 {
	return 0
}

func (m *MessageHeader) VersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.VersionSinceVersion()
}

func (*MessageHeader) VersionDeprecated() uint16 {
	return 0
}

func (*MessageHeader) NumGroupsMinValue() uint16 {
	return 0
}

func (*MessageHeader) NumGroupsMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MessageHeader) NumGroupsNullValue() uint16 {
	return math.MaxUint16
}

func (*MessageHeader) NumGroupsSinceVersion() uint16 {
	return 0
}

func (m *MessageHeader) NumGroupsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NumGroupsSinceVersion()
}

func (*MessageHeader) NumGroupsDeprecated() uint16 {
	return 0
}

func (*MessageHeader) NumVarDataFieldsMinValue() uint16 {
	return 0
}

func (*MessageHeader) NumVarDataFieldsMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MessageHeader) NumVarDataFieldsNullValue() uint16 {
	return math.MaxUint16
}

func (*MessageHeader) NumVarDataFieldsSinceVersion() uint16 {
	return 0
}

func (m *MessageHeader) NumVarDataFieldsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NumVarDataFieldsSinceVersion()
}

func (*MessageHeader) NumVarDataFieldsDeprecated() uint16 {
	return 0
}
