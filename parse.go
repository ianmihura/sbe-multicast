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

	obj := sbeParser(buff, true)
	obj.PPrint(0)
}

func sbeParser(buff []byte, doPrintFrameHeader bool) deribit_multicast.SbeStdMessage {
	// TODO increase buffer size of SbeMarshalling for parsing (excluding variable length)

	// TODO reader smaller : more efficient
	r := bytes.NewReader(buff)
	sbe_m := deribit_multicast.NewSbeGoMarshaller()

	frame := deribit_multicast.FrameHeader{}
	err := frame.Decode(sbe_m, r, 1)
	if err != nil {
		log.Fatal(err)
	}
	if doPrintFrameHeader {
		frame.PPrint(0)
	}

	// TODO refactor header outside of packet
	// TODO switch on message type, will be in the header

	// header := deribit_multicast.MessageHeader{}
	// err := header.Decode(sbe_m, r, 1) // TODO actingversion?
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(">>>> DECODED : MessageHeader", header)

	rfq := deribit_multicast.Rfq{}
	// rfq decodes its own header
	err = rfq.Decode(sbe_m, r, 1, frame.PacketLength, true) // TODO version before reading?
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(">>>> DECODED : RFQ", rfq)

	return &rfq
}
