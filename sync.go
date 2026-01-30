package main

import (
	"fmt"
	"os"
	"sync/atomic"

	"github.com/ianmihura/sbe-multicast/stdmsg"
)

var LOOK_AHEAD = 5

func SyncWorkers(rBuffs *RingBuffers, killCh chan<- os.Signal) {
	order := uint32(0)
	// rcv := 0
	// start := time.Now()

	readers := *InitRingBuffersReaders()
	fmt.Println(readers[2])
	chanIdx := 0

	for {
		// We do one channel at a time
		// This is unrealistic for a real-world scenario:
		// You'll probably have different processes managing each
		// channel, based on the info provided
		cChannelId := ChannelIds[chanIdx]
		// Move to next channel (prepare next loop)
		chanIdx = (chanIdx + 1) % len(ChannelIds)

		msgs := (*rBuffs)[cChannelId][readers[cChannelId]]
		if msgs == nil {
			// Gap Detection: If current is nil but future exists, skip current
			futureIdx := (readers[cChannelId] + LOOK_AHEAD) & (RingBufferSize - 1)
			if (*rBuffs)[cChannelId][futureIdx] != nil {
				fmt.Printf("Gap detected on channel %d at seq %d, skipping...\n", cChannelId, readers[cChannelId])
				readers[cChannelId] = (readers[cChannelId] + 1) & (RingBufferSize - 1)
			}
			continue
		}

		// TODO monitor messages processed per sec
		// if IsM {
		// 	rcv++
		// 	PrintNetworkMonitor(rcv, &start, "Processed")
		// }
		if IsP {
			fmt.Println()
			for _, msg := range msgs {
				msg.PPrint(0)
			}
		}

		if IsV {
			for _, msg := range msgs {
				s, ok := msg.(*stdmsg.PriceIndex)
				if ok {
					fmt.Println(s.Header.SequenceNumber, order)
					atomic.AddUint32(&order, 1)
				}
			}
		}

		// Set to nil when we done reading it
		(*rBuffs)[cChannelId][readers[cChannelId]] = nil
		readers[cChannelId] = (readers[cChannelId] + 1) & (RingBufferSize - 1)
	}
}

// Each channel has its own ring buffer,
// each ring buffer holds a slice of *stdmsg.StdMessage
// because each frame received has one or more messages.
// Total of 26*2 +1 = 53 channels (id=0-26, 101-126).
//
// RingBuffers Usage:
//   - RingBuffer[channelId][seqNum] == nil
//   - RingBuffer[channelId][seqNum][messageIdx]
type RingBuffers map[uint16][][]stdmsg.StdMessage

const RingBufferSize = _1KB / 8 // 128 = 1024/8 (8 byte size pointer, keep each RingBuff @1KB)

func InitRingBuffers() *RingBuffers {
	buffs := make(RingBuffers, len(ChannelIds))
	for _, id := range ChannelIds {
		buffs[id] = make([]([]stdmsg.StdMessage), RingBufferSize)
	}
	return &buffs
}

// RingBuffersReaders is compound index of frames,
// a map of channelId to next expected seqNum to receive
// Total of 26*2 +1 = 53 channels (id=0-26, 101-126).
type RingBuffersReaders map[uint16]int

func InitRingBuffersReaders() *RingBuffersReaders {
	readers := make(RingBuffersReaders, len(ChannelIds))
	for _, id := range ChannelIds {
		readers[id] = 0
	}
	return &readers
}

// Valid range of channels (0-26 and 101-126)
var ChannelIds = func() []uint16 {
	ids := make([]uint16, 0, 53)
	for i := uint16(0); i <= 26; i++ {
		ids = append(ids, i)
	}
	for i := uint16(101); i <= 126; i++ {
		ids = append(ids, i)
	}
	return ids
}()
