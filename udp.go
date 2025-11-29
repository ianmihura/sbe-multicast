package main

import (
	"encoding/hex"
	"log"
	"net"
	"sync"
	"syscall"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const IFACE = "wlan0"
const LO_ADDR = "127.0.0.1"

// Need to init the conn with a specific sockopt for iface=lo
func getConnLO() (conn *net.UDPConn) {
	localAddr, err := net.ResolveUDPAddr("udp4", LO_ADDR+":0")
	if err != nil {
		log.Fatal("error in udp:", err)
	}

	conn, err = net.ListenUDP("udp4", localAddr)
	if err != nil {
		log.Fatal("error in udp:", err)
	}

	rawConn, err := conn.SyscallConn()
	if err != nil {
		log.Fatal("error in udp:", err)
	}

	// mark sender to receive its own packets
	if rawConn.Control(func(fd uintptr) {
		err := syscall.SetsockoptInt(int(fd), syscall.IPPROTO_IP, syscall.IP_MULTICAST_LOOP, 1)
		if err != nil {
			log.Fatal("error in udp config:", err)
		}
	}) != nil {
		log.Fatal("error in udp config:", err)
	}

	return conn
}

// Pings an number (even increasing) to a udp addr
func PingUDPLoopback(addr_ string) {
	if IFACE != "lo" {
		log.Fatalln("error in udp: const IFACE != `lo`")
	}

	addr, err := net.ResolveUDPAddr("udp4", addr_)
	if err != nil {
		log.Fatal("error in udp:", err)
	}

	conn := getConnLO()
	defer conn.Close()

	var i int32 = 0
	for {
		_, err := conn.WriteToUDP([]byte(string(i)), addr)
		if err != nil {
			log.Fatal("error in udp:", err)
		}

		i += 1
	}
}

// Listens for incoming multicast messages from an addr, uses handler_ callback
func ListenUDPSingle(addr_ string, handler_ func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp", addr_)
	if err != nil {
		log.Fatal("error in udp:", err)
	}

	if_addr, err := net.InterfaceByName(IFACE)
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
		return make([]byte, _4KB)
	},
}

// Listens for incoming multicast messages from an addr, sends messages via dataCh.
// Can also dump hex payload to stdout if isLogging
func ListenUDPFast(addr_ string, dataCh chan<- []byte, isLogging bool) {
	addr, err := net.ResolveUDPAddr("udp4", addr_)
	if err != nil {
		log.Fatal("error in udp listener:", err)
	}

	if_addr, err := net.InterfaceByName(IFACE)
	if err != nil {
		log.Fatal("error in udp listener:", err)
	}

	conn, err := net.ListenMulticastUDP("udp4", if_addr, addr)
	if err != nil {
		log.Fatal("error in udp: listener", err)
	}

	conn.SetReadBuffer(_64KB)

	log.Println("Listening on", if_addr, "from", addr)
	for {
		buff := buffPool.Get().([]byte)[:_4KB]

		nBytes, src, err := conn.ReadFromUDP(buff)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		} else if isLogging {
			log.Println(nBytes, "received from addr", src)
			log.Printf("payload dump:\n%s", hex.Dump(buff[:nBytes]))
		}

		dataCh <- buff[:nBytes]
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

	log.Println("Sending on", addr)
	for {
		for _, packet := range packets {
			conn.Write(packet.ApplicationLayer().Payload())
		}
	}
}
