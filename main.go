package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/ianmihura/sbe-multicast/stdmsg"
)

const _4KB = 4096
const _64KB = 65_536

const MC_GROUP = "239.222.222.2"
const MC_PORT = "6200"
const FILE = "./pcaps/sample_capture_v1_6.pcapng"
const DATA_CHAN_CAP = 100 // TODO maybe use un-buffered chan?

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	nProc := runtime.NumCPU()
	runtime.GOMAXPROCS(nProc)

	addr := MC_GROUP + ":" + MC_PORT

	// Using single thread we can parse each packet right after reception
	// ListenUDPSingle(addr, ParseSingle)
	// Otherwise we spinup a single fast UDP listener that delegates
	//   parsing to parser workers

	// dataCh will grab incoming packets from socket
	//   and carry each packet (message) from Listener to Parser,
	dataCh := make(chan []byte, DATA_CHAN_CAP)
	go ListenUDPFast(addr, dataCh, false)

	// syncCh will grab finished work of each worker
	//   and send it to be executed in-line
	syncCh := make(chan *stdmsg.StdMessage, nProc*2)

	// We spinup n workers (as NumCPU)
	for range nProc {
		go ParseWorker(dataCh, syncCh)
	}

	// Sync up the work as we receive it
	go SyncWorkers(syncCh)

	// Replay packets
	time.Sleep(time.Second) // sleep 1sec to allow listener & workers to spinup
	// go PingUDPLoopback(addr)
	go ReplayUDP(FILE, addr)

	// Keep app alive
	c := make(chan os.Signal, 1)
	s := <-c
	log.Printf("Received signal '%v', halting program\n", s)
}

// https://jewelhuq.medium.com/mastering-high-performance-tcp-udp-socket-programming-in-go-996dc85f5de1
// https://stackoverflow.com/questions/60337662/how-to-maximise-udp-packets-per-second-with-go
// https://blog.cloudflare.com/how-to-receive-a-million-packets/
// https://tungdam.medium.com/linux-network-ring-buffers-cea7ead0b8e8
// https://ntk148v.github.io/posts/linux-network-performance-ultimate-guide/
// https://balodeamit.blogspot.com/2013/10/receive-side-scaling-and-receive-packet.html
// https://docs.redhat.com/en/documentation/red_hat_enterprise_linux/10/html/network_troubleshooting_and_performance_tuning/tuning-network-adapter-settings
// https://blog.packagecloud.io/monitoring-tuning-linux-networking-stack-receiving-data/
