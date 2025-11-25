// Generated SBE (Simple Binary Encoding) message codec

package deribit_multicast

import (
	"fmt"
	"io"
	"math"
)

type MessageHeader struct {
	BlockLength      uint16
	TemplateId       uint16
	SchemaId         uint16
	Version          uint16
	NumGroups        uint16
	NumVarDataFields uint16
}

func (h *MessageHeader) PPrint(i int) {
	PPrintlnInd(i, "Object: Header")
	PPrintlnInd(i+2, "BlockLength:", h.BlockLength)
	PPrintlnInd(i+2, "TemplateId:", h.getTemplateName())
	PPrintlnInd(i+2, "SchemaId:", h.SchemaId)
	PPrintlnInd(i+2, "Version:", h.Version)
	PPrintlnInd(i+2, "NumGroups:", h.NumGroups)
	PPrintlnInd(i+2, "NumVarDataFields:", h.NumVarDataFields)
}

func (h *MessageHeader) getTemplateName() string {
	switch h.TemplateId {
	case 1009:
		return "RFQ (1009)"
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

func (m *MessageHeader) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !m.BlockLengthInActingVersion(actingVersion) {
		m.BlockLength = m.BlockLengthNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.BlockLength); err != nil {
			return err
		}
	}
	if !m.TemplateIdInActingVersion(actingVersion) {
		m.TemplateId = m.TemplateIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.TemplateId); err != nil {
			return err
		}
	}
	if !m.SchemaIdInActingVersion(actingVersion) {
		m.SchemaId = m.SchemaIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.SchemaId); err != nil {
			return err
		}
	}
	if !m.VersionInActingVersion(actingVersion) {
		m.Version = m.VersionNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.Version); err != nil {
			return err
		}
	}
	if !m.NumGroupsInActingVersion(actingVersion) {
		m.NumGroups = m.NumGroupsNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.NumGroups); err != nil {
			return err
		}
	}
	if !m.NumVarDataFieldsInActingVersion(actingVersion) {
		m.NumVarDataFields = m.NumVarDataFieldsNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.NumVarDataFields); err != nil {
			return err
		}
	}
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
