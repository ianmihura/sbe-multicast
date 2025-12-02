package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/ianmihura/sbe-multicast/stdmsg"

	_ "net/http/pprof"
)

const _4KB = 4096
const _64KB = 65_536

const MC_GROUP = "239.222.222.2"
const MC_PORT = "6200"
const FILE = "./pcaps/sample_capture_v1_6.pcapng"
const DATA_CHAN_CAP = 100 // TODO maybe use un-buffered chan?

var Mode *string
var Iface *string
var IsLoop bool = false
var IsP bool = false
var IsH bool = false
var IsM bool = false

func main() {
	// TODO choose sample file
	Mode = flag.String("mode", "sample", "Replay mode.\n  Options: [ping, sample].")
	Iface = flag.String("iface", "wlan0", "Network interface.\n  Check available ifaces with `ip link`")
	isLoop_ := flag.Bool("l", false, "Loop: inifinite loop for pkt replay. Otherwise:\n  ping: 10 pkts.\n  sample: replay sample once")
	isp_ := flag.Bool("p", false, "Pretty-Print parsed SBE structs")
	ish_ := flag.Bool("h", false, "Hex dump received network pkts\n  (if not looping, app may end before printing full hex dump)")
	ism_ := flag.Bool("m", false, "Monitoring network pkts (sent & received)")
	flag.Parse()
	IsLoop = *isLoop_
	IsP = *isp_ && (*Mode == "sample")
	IsH = *ish_
	IsM = *ism_

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	nProc := runtime.NumCPU()
	runtime.GOMAXPROCS(nProc)
	addr := MC_GROUP + ":" + MC_PORT

	// kill signal - quit gracefully
	killCh := make(chan os.Signal, 1)

	// Using single thread we can parse each packet right after reception
	// ListenUDPSingle(addr, ParseSingle)
	// Otherwise we spinup a single fast UDP listener that delegates
	//   parsing to parser workers

	// dataCh will grab incoming packets from socket
	//   and carry each packet (message) from Listener to Parser
	dataCh := make(chan []byte, DATA_CHAN_CAP)
	go ListenUDPFast(addr, dataCh)

	// syncCh will grab finished work of each worker
	//   and send it to be executed in-line
	syncCh := make(chan *stdmsg.StdMessage, nProc*2)

	// We spinup n workers (as NumCPU)
	for range nProc {
		go ParseWorker(dataCh, syncCh)
	}

	// Sync up the work as we receive it
	go SyncWorkers(syncCh, killCh)

	// Send packets
	if !IsLoop {
		time.Sleep(time.Second) // sleep 1sec to allow listener & workers to spinup
	}
	if *Mode == "ping" {
		go PingUDP(addr, killCh)
	} else {
		go ReplayUDP(FILE, addr, killCh)
	}

	// For pperf
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Keep app alive
	s := <-killCh
	log.Printf("Received signal '%v', halting program\n", s)
}
