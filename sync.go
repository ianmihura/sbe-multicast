package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/ianmihura/sbe-multicast/stdmsg"
)

func SyncWorkers(syncCh <-chan *stdmsg.StdMessage, killCh chan<- os.Signal) {
	// TODO ordering the messages
	// TODO checking for missing packets
	// TODO monitoring performance
	// TODO saving data for replay (eg. protobuf)

	var order uint32 = 0
	rcv := 0
	last := time.Now()

	for {
		msg, ok := <-syncCh
		if !ok {
			killCh <- os.Kill
		}

		if IsM {
			rcv++
			PrintNetworkMonitor(rcv, &last, "Processed")
		}
		if IsP {
			fmt.Println()
			(*msg).PPrint(0)
		}

		if IsV {
			s, ok := (*msg).(*stdmsg.PriceIndex)
			if ok {
				fmt.Println(s.Header.SequenceNumber, order)
				// if s.Header.SequenceNumber != order {
				// fmt.Println(s.Header.SequenceNumber, order)
				// }
				atomic.AddUint32(&order, 1)
			}
		}
	}
}
