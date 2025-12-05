package main

import (
	"encoding/hex"
	"log"
	"net"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const LO_ADDR = "127.0.0.1"

// Pings an number (even increasing) to a udp addr
func PingUDP(addr_ string, killCh chan<- os.Signal) {
	addr, err := net.ResolveUDPAddr("udp4", addr_)
	if err != nil {
		log.Fatal("error in udp:", err)
	}

	conn := getConn(addr)
	defer conn.Close()

	i := 0
	last := time.Now()
	for {
		if *Iface == "lo" {
			_, err = conn.WriteToUDP([]byte(string(i)), addr)
		} else {
			_, err = conn.Write([]byte(string(i)))
		}
		if err != nil {
			log.Fatal("error in udp:", err)
		}

		i++

		if IsM {
			PrintNetworkMonitor(i, &last, "Sent")
		}

		if !IsLoop && i >= 9 {
			time.Sleep(time.Second / 10) // wait for all parse workers to finish their work
			killCh <- os.Kill
		}
	}
}

var buffPool = sync.Pool{
	New: func() any {
		return make([]byte, _1KB)
	},
}

// Listens for incoming multicast messages from an addr, sends messages via dataCh.
// Can also dump hex payload to stdout if isLogging
func ListenUDPFast(addr_ string, dataCh chan<- []byte) {
	addr, err := net.ResolveUDPAddr("udp4", addr_)
	if err != nil {
		log.Fatal("error in udp listener:", err)
	}

	if_addr, err := net.InterfaceByName(*Iface)
	if err != nil {
		log.Fatal("error in udp listener:", err)
	}

	conn, err := net.ListenMulticastUDP("udp4", if_addr, addr)
	if err != nil {
		log.Fatal("error in udp: listener", err)
	}

	conn.SetReadBuffer(_1KB * 128)

	log.Println("Listening on", if_addr, "from", addr)
	for {
		buff := buffPool.Get().([]byte)[:_1KB]

		nBytes, src, err := conn.ReadFromUDP(buff)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		} else if IsH {
			log.Println(nBytes, "received from addr", src)
			log.Printf("payload dump:\n%s", hex.Dump(buff[:nBytes]))
		}

		dataCh <- buff[:nBytes]
	}
}

// Replays pcap or pcapng content to a specific group
func ReplayUDP(file string, addr_ string, killCh chan<- os.Signal) {
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

	conn := getConn(addr)
	defer conn.Close()

	log.Println("Sending on", addr)
	pktsent := 0
	last := time.Now()
	for {
		for _, packet := range packets {
			var err error
			if *Iface == "lo" {
				_, err = conn.WriteToUDP(packet.ApplicationLayer().Payload(), addr)
			} else {
				_, err = conn.Write(packet.ApplicationLayer().Payload())
			}
			if err != nil {
				log.Fatal("error in udp replay:", err)
			}

			if IsM {
				pktsent++
				PrintNetworkMonitor(pktsent, &last, "Sent")
			}
		}
		if !IsLoop {
			time.Sleep(time.Second / 10) // wait for all parse workers to finish their work
			killCh <- os.Kill
		}
	}
}

// Returns a udp connection to a specified addr, based on the Iface
func getConn(addr *net.UDPAddr) (conn *net.UDPConn) {
	if *Iface == "lo" {

		// Need to init the conn with a specific sockopt for iface=lo
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
	} else {
		var err error
		conn, err = net.DialUDP("udp4", nil, addr)
		if err != nil {
			log.Fatal("error in udp replay:", err)
		}
	}

	return conn
}
