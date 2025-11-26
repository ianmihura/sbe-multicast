package main

import (
	"bytes"
	"encoding/hex"
	"log"
	"net"

	"github.com/ianmihura/sbe-multicast/deribit_multicast"
)

func Parser(src *net.UDPAddr, nBytes int, buff []byte) {
	log.Println(nBytes, "received from addr", src)
	log.Printf("payload dump:\n%s", hex.Dump(buff[:nBytes]))

	msg := sbeParser(buff, true)
	msg.PPrint(0)
}

func sbeParser(buff []byte, doPrintFrameHeader bool) deribit_multicast.SbeStdMessage {
	// TODO reader smaller : more efficient
	r := bytes.NewReader(buff)
	sbe_m := deribit_multicast.NewSbeGoMarshaller()

	frame := deribit_multicast.FrameHeader{}
	err := frame.Decode(sbe_m, r)
	if err != nil {
		log.Fatal(err)
	}
	if doPrintFrameHeader {
		frame.PPrint(0)
	}

	header := deribit_multicast.MessageHeader{}
	err = header.Decode(sbe_m, r)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := header.GetConcreteMessage()
	if err != nil {
		log.Fatal(err)
	}

	err = msg.Decode(sbe_m, r, header.Version, frame.PacketLength, true)
	if err != nil {
		log.Fatal(err)
	}

	return msg
}
