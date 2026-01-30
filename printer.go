package main

import (
	"log"
	"os"
	"time"
)

func PrintNetworkMonitor(pkts int, start *time.Time, action string) {
	var every int
	if IsLoop {
		every = 10_000
	} else {
		every = 1_013 // total for this capture is 5066 (5065 is close enough)
	}

	if pkts%every == 0 {
		elapsed := time.Since(*start).Seconds()
		pps := int(float64(pkts) / elapsed)

		log.Printf("Pkts %s: %d | PPS: %d", action, pkts, pps)
	}
}

func VPrint(killCh chan<- os.Signal) {
	// for {
	// 	time.Sleep(time.Second * 2)
	// 	log.Println(Freq)
	// 	log.Println()
	// 	killCh <- os.Kill
	// }
}
