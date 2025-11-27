package main

import (
	"github.com/ianmihura/sbe-multicast/stdmsg"
)

func SyncWorkers(syncCh <-chan *stdmsg.StdMessage) {
	// Here we may consider:
	// - ordering the messages
	// - checking for missing packets
	// - monitoring performance
	// - saving data for replay (eg. protobuf)

	for msg := range syncCh {
		(*msg).PPrint(0)
	}
}
