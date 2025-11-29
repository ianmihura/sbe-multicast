package main

import (
	"log"
	"time"

	"github.com/ianmihura/sbe-multicast/stdmsg"
)

func SyncWorkers(syncCh <-chan *stdmsg.StdMessage) {
	// Here we may consider:
	// - ordering the messages
	// - checking for missing packets
	// - monitoring performance
	// - saving data for replay (eg. protobuf)

	rcv := 0
	totrcv := 0
	last := time.Now()

	for {
		<-syncCh
		rcv++

		if rcv >= 10_000 {
			now := time.Now()
			elapsed := now.Sub(last)
			pps := int(float64(rcv) / elapsed.Seconds())

			log.Printf("Total Rcv: %d | PPS: %d", totrcv, pps)

			totrcv += rcv
			rcv = 0
			last = time.Now()
		}
	}

	// for msg := range syncCh {
	// 	// switch (*msg).(type) {
	// 	// case *stdmsg.Book:
	// 	fmt.Println()
	// 	(*msg).PPrint(0)
	// 	// default:
	// 	// 	// Handle other message types or ignore
	// 	// }
	// }
}
