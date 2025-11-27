package main

import (
	"bytes"
	"log"
	"net"
	"sync"

	"github.com/ianmihura/sbe-multicast/deribit_multicast"
	"github.com/ianmihura/sbe-multicast/stdmsg"
)

var coderPool = sync.Pool{
	New: func() any {
		return stdmsg.NewEmptyCoder()
	},
}

func ParseWorker(dataCh <-chan []byte, syncCh chan<- *stdmsg.StdMessage) {
	for data := range dataCh {
		msg := stdParser(data)
		syncCh <- &msg
	}
}

func stdParser(data []byte) stdmsg.StdMessage {
	c := coderPool.Get().(stdmsg.Coder)
	c.SetBuffer(&data)
	defer coderPool.Put(&c)

	// We can return data to dataCh once we finish using Coder
	defer buffPool.Put(&data)

	frame := stdmsg.FrameHeader{}
	frame.Decode(&c)
	// frame.PPrint(0)
	// TODO send frame (seq num) via msg to sync

	header := stdmsg.MessageHeader{}
	header.Decode(&c)

	msg, _ := header.GetConcreteMessage()
	msg.Decode(&c)

	return msg
}

func ParseSingle(src *net.UDPAddr, nBytes int, buff []byte) {
	msg := sbeParser(buff, true)
	msg.PPrint(0)
}

func sbeParser(buff []byte, isLoggingFrameHeader bool) deribit_multicast.SbeStdMessage {
	// TODO reader smaller : more efficient
	r := bytes.NewReader(buff)
	// TODO frame is always 8, make this a constant
	// TODO i should instanciate this sbe_m buffer outside
	sbe_m := deribit_multicast.NewSbeGoMarshaller(8, _8KB)

	frame := deribit_multicast.FrameHeader{}
	err := frame.Decode(sbe_m, r)
	if err != nil {
		log.Fatal(err)
	}
	if isLoggingFrameHeader {
		frame.PPrint(0)
	}

	// TODO resize or send the size of buff to read to the readatleast
	sbe_m.Resize(12) // message header is always 12
	header := deribit_multicast.MessageHeader{}
	err = header.Decode(sbe_m, r)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := header.GetConcreteMessage()
	if err != nil {
		log.Fatal(err)
	}

	sbe_m.Resize(int(header.BlockLength)) // rfq is always size 22, anyway this exists as header.BlockLength
	// fmt.Println(sbe_m)
	err = msg.Decode(sbe_m, r, header.Version, frame.PacketLength, true)
	if err != nil {
		log.Fatal(err)
	}

	return msg
}
