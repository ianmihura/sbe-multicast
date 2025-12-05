package main

import (
	"log"
	"os"
	"time"
)

func PrintNetworkMonitor(pkts int, last *time.Time, action string) {
	var every int
	if IsLoop {
		every = 10_000
	} else {
		every = 1
	}

	if pkts%every == 0 {
		elapsed := time.Since(*last).Seconds()
		pps := int(float64(pkts) / elapsed)

		// TODO monitor mbps
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
