package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/ianmihura/sbe-multicast/stdmsg"

	"net/http"
	"net/http/pprof"
)

const _1KB = 1024

const MC_GROUP = "239.222.222.2"
const MC_PORT = "6200"

// const FILE = "./pcaps/sample_capture_v1_6.pcapng"

const FILE = "./pcaps/price_index.pcapng"
const DATA_CHAN_CAP = 1000
const SYNC_CHAN_CAP = 1000

// var Freq = make([]int32, SYNC_CHAN_CAP+1)

var Mode *string
var Iface *string
var IsLoop bool = false
var IsP bool = false
var IsH bool = false
var IsM bool = false
var IsV bool = false
var PProf bool = false

func main() {
	argparse()
	if PProf {
		go pprofServer()
	}

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	nProc := runtime.NumCPU()
	runtime.GOMAXPROCS(nProc)
	addr := MC_GROUP + ":" + MC_PORT

	// kill signal - quit gracefully
	killCh := make(chan os.Signal, 1)

	// `dataCh` will grab incoming packets from socket
	//   and carry each packet (message) from Listener to Parser
	dataCh := make(chan []byte, DATA_CHAN_CAP)
	go ListenUDPFast(addr, dataCh)

	// `syncCh` will grab finished work of each worker
	//   and send it to be executed in-line
	syncCh := make(chan *stdmsg.StdMessage, SYNC_CHAN_CAP)

	// We spinup n workers (as NumCPU)
	for i := range nProc {
		go ParseWorker(dataCh, syncCh, uint32(i))
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

	if IsV {
		go VPrint(killCh)
	}

	// Keep app alive
	s := <-killCh
	log.Printf("Received signal '%v', halting program\n", s)
}

func argparse() {
	// TODO choose sample file
	Mode = flag.String("mode", "sample", "Replay mode.\n  Options: [ping, sample].")
	Iface = flag.String("iface", "wlan0", "Network interface.\n  Check available ifaces with `ip link`")
	isLoop_ := flag.Bool("l", false, "Loop: inifinite loop for pkt replay. Otherwise:\n  ping: 10 pkts.\n  sample: replay sample once")
	isp_ := flag.Bool("p", false, "Pretty-Print parsed SBE structs\n  (if not looping, app may end before printing full hex dump)")
	ish_ := flag.Bool("h", false, "Hex dump received network pkts\n  (if not looping, app may end before printing full hex dump)")
	ism_ := flag.Bool("m", false, "Monitoring network pkts (sent & received)")
	isv_ := flag.Bool("v", false, "Verbose. Unstructured monitoring.")
	pprof_ := flag.Bool("pprof", false, "Serve pprof standard server on localhost:8080")

	flag.Parse()

	IsLoop = *isLoop_
	IsP = *isp_ && (*Mode == "sample")
	IsH = *ish_
	IsM = *ism_
	IsV = *isv_
	PProf = *pprof_
}

func pprofServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile) // CPU

	go http.ListenAndServe(":8080", mux)

	log.Println("pprof server up in localhost:8080/debug/")
}
