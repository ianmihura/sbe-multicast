package main

import (
	"log"
	"net"
	"sync"

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
	c := coderPool.Get().(*stdmsg.Coder)
	c.SetBuffer(&data)
	defer coderPool.Put(c)

	// We can return data to dataCh once we finish using Coder
	defer buffPool.Put(&data)

	frame := stdmsg.FrameHeader{}
	frame.Decode(c)
	// frame.PPrint(0)
	// TODO send frame (seq num) via msg to sync

	header := stdmsg.MessageHeader{}
	header.Decode(c)

	msg, err := header.GetConcreteMessage()
	if err != nil {
		log.Fatal("error in stdParser:", err)
	}
	msg.Decode(c)

	return msg
}

func ParseSingle(src *net.UDPAddr, nBytes int, buff []byte) {
	msg := stdParser(buff)
	msg.PPrint(0)
}
