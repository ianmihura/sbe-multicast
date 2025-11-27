package main

import (
	"encoding/hex"
	"log"
	"net"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Pings an number (even increasing) to a udp addr
func PingUDP(addr_ string) {
	addr, err := net.ResolveUDPAddr("udp4", addr_)
	if err != nil {
		log.Fatal("error in udp:", err)
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal("error in udp:", err)
	}

	var i int32 = 0
	for {
		nBytes, err := conn.Write([]byte(string(i)))
		if err != nil {
			log.Fatal("error in udp:", err)
		}

		log.Println(nBytes, "sent to addr", addr)
		time.Sleep(time.Second)
		i += 1
	}
}

// Listens for incoming multicast messages from an addr, uses handler_ callback
func ListenUDPSingle(addr_ string, handler_ func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp", addr_)
	if err != nil {
		log.Fatal("error in udp:", err)
	}

	if_addr, err := net.InterfaceByName("wlan0")
	if err != nil {
		log.Fatal("error in udp:", err)
	}

	conn, err := net.ListenMulticastUDP("udp", if_addr, addr)
	if err != nil {
		log.Fatal("error in udp:", err)
	}

	conn.SetReadBuffer(_64KB)
	buff := make([]byte, _64KB)

	log.Println("Listening on", if_addr, "from", addr)
	for {
		nBytes, src, err := conn.ReadFromUDP(buff)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}

		log.Println(nBytes, "received from addr", src)
		log.Printf("payload dump:\n%s", hex.Dump(buff[:nBytes]))

		handler_(src, nBytes, buff)
	}
}

var buffPool = sync.Pool{
	New: func() any {
		buff := make([]byte, _8KB) // TODO check max msg size
		return &buff
	},
}

// Listens for incoming multicast messages from an addr, sends messages via dataCh
func ListenUDPFast(addr_ string, dataCh chan<- []byte, isLogging bool) {
	addr, err := net.ResolveUDPAddr("udp", addr_)
	if err != nil {
		log.Fatal("error in udp listener:", err)
	}

	if_addr, err := net.InterfaceByName("wlan0")
	if err != nil {
		log.Fatal("error in udp listener:", err)
	}

	conn, err := net.ListenMulticastUDP("udp", if_addr, addr)
	if err != nil {
		log.Fatal("error in udp: listener", err)
	}

	// TODO measure biggest message size
	conn.SetReadBuffer(_64KB)

	log.Println("Listening on", if_addr, "from", addr)
	for {
		buff := buffPool.Get().(*[]byte)

		nBytes, src, err := conn.ReadFromUDP(*buff)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		} else if isLogging {
			log.Println(nBytes, "received from addr", src)
			log.Printf("payload dump:\n%s", hex.Dump((*buff)[:nBytes]))
		}

		dataCh <- (*buff)[:nBytes]
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
		log.Fatal("error in udp replay:", err)
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal("error in udp replay:", err)
	}
	defer conn.Close()

	log.Println("Sending on", addr_)
	// for {
	for _, packet := range packets {
		// time.Sleep(time.Second)
		conn.Write(packet.ApplicationLayer().Payload())
	}
	// log.Println(">>> eof, replaying caputre")
	// }
}
