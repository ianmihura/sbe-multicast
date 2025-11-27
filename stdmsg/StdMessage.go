package stdmsg

import (
	"encoding/binary"
	"fmt"
	"log"
)

type Coder struct {
	order binary.ByteOrder
	buff  []byte
	off   int
}

func NewEmptyCoder() Coder {
	return Coder{
		order: binary.LittleEndian,
		buff:  []byte{},
		off:   0,
	}
}

func NewCoder(buff []byte) Coder {
	return Coder{
		order: binary.LittleEndian,
		buff:  buff,
		off:   0,
	}
}

func (c *Coder) SetBuffer(buff *[]byte) {
	c.buff = *buff
}

func (c *Coder) Decode(data any) {
	n, err := binary.Decode(c.buff[c.off:], c.order, data)
	if err != nil {
		log.Fatal(err)
	}
	c.off += n
}

type StdMessage interface {
	Decode(c *Coder)
	PPrint(indent int)
}

func PPrintlnInd(indent int, args ...any) {
	is := ""
	for range indent {
		is += " "
	}

	fmt.Print(is)
	fmt.Println(args...)
}
