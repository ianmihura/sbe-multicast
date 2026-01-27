package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ianmihura/sbe-multicast/stdmsg"
)

var coderPool = sync.Pool{
	New: func() any {
		return stdmsg.NewEmptyCoder()
	},
}

func ParseWorker(dataCh <-chan []byte, syncCh chan<- *stdmsg.StdMessage, start *time.Time, rcv *int32) {
	var rcv_ int32
	for data := range dataCh {
		if *Mode == "ping" {
			syncCh <- nil

		} else {
			c := coderPool.Get().(*stdmsg.Coder)
			c.SetBuffer(&data)
			c.ResetOffset()
			defer coderPool.Put(c)

			// We can return data to dataCh pool once we finish using Coder
			defer buffPool.Put(data)

			frame := stdmsg.FrameHeader{}
			frame.Decode(c)

			for uint16(c.GetOffset()) < frame.PacketLength {
				header := stdmsg.MessageHeader{
					SequenceNumber: frame.SequenceNumber,
					ChannelId:      frame.ChannelId,
				}
				header.Decode(c)

				msg, err := header.GetConcreteMessage()
				if err != nil {
					log.Fatal(err)
				}
				msg.Decode(c)
				syncCh <- &msg
			}

			if IsM {
				rcv_ = atomic.AddInt32(rcv, 1)
				go PrintNetworkMonitor(int(rcv_), start, "Processed")
			}
		}
	}
}
