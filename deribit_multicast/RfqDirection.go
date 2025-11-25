// Generated SBE (Simple Binary Encoding) message codec

package deribit_multicast

import (
	"fmt"
	"io"
	"reflect"
)

type RfqDirectionEnum uint8
type RfqDirectionValues struct {
	Buy          RfqDirectionEnum
	Sell         RfqDirectionEnum
	No_direction RfqDirectionEnum
	NullValue    RfqDirectionEnum
}

var RfqDirection = RfqDirectionValues{0, 1, 2, 255}

func (rd *RfqDirectionEnum) GetPPrint() string {
	switch *rd {
	case RfqDirection.Buy:
		return "buy"
	case RfqDirection.Sell:
		return "sell"
	case RfqDirection.No_direction:
		return "no direction"
	default:
		return "null"
	}
}

func (r RfqDirectionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(r)); err != nil {
		return err
	}
	return nil
}

func (r *RfqDirectionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(r)); err != nil {
		return err
	}
	return nil
}

func (r RfqDirectionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(RfqDirection)
	for idx := 0; idx < value.NumField(); idx++ {
		if r == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("range check failed on RfqDirection, unknown enumeration value %d", r)
}

func (*RfqDirectionEnum) EncodedLength() int64 {
	return 1
}

func (*RfqDirectionEnum) buySinceVersion() uint16 {
	return 0
}

func (r *RfqDirectionEnum) buyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.buySinceVersion()
}

func (*RfqDirectionEnum) buyDeprecated() uint16 {
	return 0
}

func (*RfqDirectionEnum) sellSinceVersion() uint16 {
	return 0
}

func (r *RfqDirectionEnum) sellInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.sellSinceVersion()
}

func (*RfqDirectionEnum) sellDeprecated() uint16 {
	return 0
}

func (*RfqDirectionEnum) no_directionSinceVersion() uint16 {
	return 0
}

func (r *RfqDirectionEnum) no_directionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.no_directionSinceVersion()
}

func (*RfqDirectionEnum) no_directionDeprecated() uint16 {
	return 0
}
