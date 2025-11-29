package stdmsg

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
)

// Decoder struct, inspired by encoding/binary : coder struct
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

// Decodes binary data from c.buff starting at c.off, into data according to c.order
// It quits if buf is too small, dumping useful data to stdout.
func (c *Coder) Decode(data any) {
	n, err := binary.Decode(c.buff[c.off:], c.order, data)
	if err != nil {
		log.Printf("%p:%d | %d len(%d)\n", &c.buff, c.off, n, len(c.buff))
		log.Println(hex.Dump(c.buff))
		log.Fatal("error in decode: ", err)
	}
	c.off += n
}

type StdMessage interface {
	Decode(c *Coder)   // Decodes with binary.Decode
	PPrint(indent int) // Pretty Print with an indent
}

// fmt.Println with `indent` number of spaces before.
func PPrintlnInd(indent int, args ...any) {
	is := ""
	for range indent {
		is += " "
	}

	fmt.Print(is)
	fmt.Println(args...)
}
