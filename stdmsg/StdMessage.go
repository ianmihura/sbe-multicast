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

func NewEmptyCoder() *Coder {
	return &Coder{
		order: binary.LittleEndian,
		buff:  []byte{},
		off:   0,
	}
}

func (c *Coder) SetBuffer(buff *[]byte) {
	c.buff = *buff
}

func (c *Coder) ResetOffset() {
	c.off = 0
}

func (c *Coder) Decode(data any) {
	n, err := binary.Decode(c.buff[c.off:], c.order, data)
	if err != nil {
		// fmt.Printf("%p:%d | %d len(%d)\n", &c.buff, c.off, n, len(c.buff))
		// fmt.Println(hex.Dump(c.buff))
		log.Fatal("error in decode: ", err)
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
