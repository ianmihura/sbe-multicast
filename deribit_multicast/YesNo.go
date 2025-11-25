// Generated SBE (Simple Binary Encoding) message codec

package deribit_multicast

import (
	"fmt"
	"io"
	"reflect"
)

type YesNoEnum uint8
type YesNoValues struct {
	No        YesNoEnum
	Yes       YesNoEnum
	NullValue YesNoEnum
}

var YesNo = YesNoValues{0, 1, 255}

func (yn *YesNoEnum) GetPPrint() string {
	switch *yn {
	case YesNo.Yes:
		return "yes"
	case YesNo.No:
		return "no"
	default:
		return "null"
	}
}

func (y YesNoEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(y)); err != nil {
		return err
	}
	return nil
}

func (y *YesNoEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(y)); err != nil {
		return err
	}
	return nil
}

func (y YesNoEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(YesNo)
	for idx := 0; idx < value.NumField(); idx++ {
		if y == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("range check failed on YesNo, unknown enumeration value %d", y)
}

func (*YesNoEnum) EncodedLength() int64 {
	return 1
}

func (*YesNoEnum) noSinceVersion() uint16 {
	return 0
}

func (y *YesNoEnum) noInActingVersion(actingVersion uint16) bool {
	return actingVersion >= y.noSinceVersion()
}

func (*YesNoEnum) noDeprecated() uint16 {
	return 0
}

func (*YesNoEnum) yesSinceVersion() uint16 {
	return 0
}

func (y *YesNoEnum) yesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= y.yesSinceVersion()
}

func (*YesNoEnum) yesDeprecated() uint16 {
	return 0
}
