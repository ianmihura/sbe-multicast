package main

import (
	"fmt"
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

func ParseWorker(dataCh <-chan []byte, rBuffs *RingBuffers, start *time.Time, rcv *int32) {
	var rcv_ int32
	for data := range dataCh {
		if *Mode == "ping" {
			continue

		} else {
			c := coderPool.Get().(*stdmsg.Coder)
			c.SetBuffer(&data)
			c.ResetOffset()
			defer coderPool.Put(c)

			// We can return data to dataCh pool once we finish using Coder
			defer buffPool.Put(data)

			frame := stdmsg.FrameHeader{}
			frame.Decode(c)

			index := frame.SequenceNumber % RingBufferSize
			if (*rBuffs)[frame.ChannelId][index] != nil {
				// repeated frame
				fmt.Println("REPEATED FRAME:")
				fmt.Println("\nOLD FRAME:")
				(*rBuffs)[frame.ChannelId][index][0].PPrint(0)

				// TODO remove this, only for testing
				header := stdmsg.MessageHeader{
					SequenceNumber: frame.SequenceNumber,
					ChannelId:      frame.ChannelId,
				}
				header.Decode(c)
				fmt.Println("\nNEW FRAME:")
				header.PPrint(0)

				continue
			}

			// Local accumulation to avoid sync worker seeing partial frames
			msgs := make([]stdmsg.StdMessage, 0, 8)
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
				msgs = append(msgs, msg)
			}

			// Atomic-like assignment of the fully populated slice header.
			(*rBuffs)[frame.ChannelId][index] = msgs

			if IsM {
				rcv_ = atomic.AddInt32(rcv, 1)
				go PrintNetworkMonitor(int(rcv_), start, "Processed")
			}
		}
	}
}
