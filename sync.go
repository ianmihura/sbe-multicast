package main

import (
	"fmt"

	"github.com/ianmihura/sbe-multicast/stdmsg"
)

func SyncWorkers(syncCh <-chan *stdmsg.StdMessage) {
	// Here we may consider:
	// - ordering the messages
	// - checking for missing packets
	// - monitoring performance
	// - saving data for replay (eg. protobuf)

	for msg := range syncCh {
		switch (*msg).(type) {
		case *stdmsg.Book:
			fmt.Println()
			(*msg).PPrint(0)
		default:
			// Handle other message types or ignore
		}
	}
}
