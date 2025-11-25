package main

import (
	"log"
	"net"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const _8KB = 8192

var pool = sync.Pool{
	New: func() any {
		return make([]byte, _8KB)
	},
}

// Pings an number (even increasing) to a udp addr
func PingUDP(addr_ string) {
	addr, err := net.ResolveUDPAddr("udp4", addr_)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	var i int32 = 0
	for {
		nBytes, err := conn.Write([]byte(string(i)))
		if err != nil {
			log.Fatal(err)
		}

		log.Println(nBytes, "sent to addr", addr)
		time.Sleep(time.Second)
		i += 1
	}
}

// Listens for incoming multicast messages from an addr, uses handler_ callback
func ListenUDP(addr_ string, handler_ func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp", addr_)
	if err != nil {
		log.Fatal(err)
	}

	if_addr, err := net.InterfaceByName("wlan0")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenMulticastUDP("udp", if_addr, addr)
	if err != nil {
		log.Fatal(err)
	}

	conn.SetReadBuffer(_8KB)

	// more efficient for concurrecny, but we'll prob not use it here
	// buff := pool.Get().([]byte)
	// defer pool.Put(&buff)
	buff := make([]byte, _8KB)

	log.Println("Listening on", if_addr, "from", addr)
	for {
		nBytes, src, err := conn.ReadFromUDP(buff)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}

		handler_(src, nBytes, buff)
	}
}

// Replays pcap or pcapng content to a specific group
func ReplayUDP(file string, addr_ string) {
	handler, err := pcap.OpenOffline(file)
	if err != nil {
		panic(err)
	}

	packets := make([]gopacket.Packet, 0)
	psource := gopacket.NewPacketSource(handler, handler.LinkType())
	for packet := range psource.Packets() {
		packets = append(packets, packet)
	}

	addr, err := net.ResolveUDPAddr("udp4", addr_)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// for {
	for _, packet := range packets {
		// time.Sleep(time.Second)
		conn.Write(packet.ApplicationLayer().Payload())
	}
	// log.Println(">>> eof, replaying caputre")
	// }
}
