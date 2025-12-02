package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ianmihura/sbe-multicast/stdmsg"
)

func SyncWorkers(syncCh <-chan *stdmsg.StdMessage, killCh chan<- os.Signal) {
	// TODO ordering the messages
	// TODO checking for missing packets
	// TODO monitoring performance
	// TODO saving data for replay (eg. protobuf)

	rcv := 0
	last := time.Now()

	for {
		msg, ok := <-syncCh
		if !ok {
			killCh <- os.Kill
		}

		rcv++

		if IsM {
			PrintNetworkMonitor(rcv, &last, "Received")
		}
		if IsP {
			fmt.Println()
			(*msg).PPrint(0)
		}
	}
}
