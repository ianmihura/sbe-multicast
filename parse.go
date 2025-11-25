package main

import (
	"bytes"
	"log"
	"net"

	"github.com/ianmihura/sbe-multicast/deribit_multicast"
)

func Parser(src *net.UDPAddr, nBytes int, buff []byte) {
	// TODO parse payload

	// TODO smaller reader to make it more efficient
	r := bytes.NewReader(buff)

	// TODO make struct for this fucking thing
	var pkgLen uint16
	var chanID uint16
	var seqNum uint32
	sbe_m := deribit_multicast.NewSbeGoMarshaller()
	sbe_m.ReadUint16(r, &pkgLen)
	sbe_m.ReadUint16(r, &chanID)
	sbe_m.ReadUint32(r, &seqNum)

	log.Println(">>>> DECODED : FIX header", pkgLen, chanID, seqNum)

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
	err := rfq.Decode(sbe_m, r, 1, 22, true) // TODO version && len before reading?
	if err != nil {
		log.Fatal(err)
	}
	log.Println(">>>> DECODED : RFQ", rfq)

	// TODO increase buffer size of SbeMarshalling for parsing (excluding variable length)

	// log.Println(nBytes, "received from addr", src)
	// log.Printf("payload dump:\n%s", hex.Dump(buff[:nBytes]))
}
