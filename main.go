package main

const MC_GROUP = "239.111.111.1"
const MC_PORT = "9999" // TODO open this port in local

func main() {
	addr := MC_GROUP + ":" + MC_PORT
	file := "./rfq.pcapng"

	// go PingUDP(addr)
	go ReplayUDP(file, addr)
	ListenUDP(addr, Parser)
}

// https://jewelhuq.medium.com/mastering-high-performance-tcp-udp-socket-programming-in-go-996dc85f5de1
