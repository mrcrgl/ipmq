package commitchain

import (
	"github.com/edsrzf/mmap-go"
)

const (
	ShardIndexFilePathTemplate = "{$.NodeHash.ShortString}/{$.SegmentHash.ShortString}/{$.Hash}.index"
	ShardDataFilePathTemplate  = "{$.NodeHash.ShortString}/{$.SegmentHash.ShortString}/{$.Hash}.data"
)

// shard describes one replica of a given segment.
// Once you write something to a segment, it should be distributed to every shard.
type shard struct {
	// hash ...
	// TODO does the hash contain parts of its segment?
	hash []byte

	node []byte

	segmentHash []byte

	// leader of the segment, if true writing is permitted, otherwise only reading
	// TODO this is not a property, it should be verified every time by the shared raft state
	leader bool

	idx index

	//
	dat data
}

type index struct {
	mm mmap.MMap
}

type block struct {
	hash Hash // will be verified at runtime

	// prev hash
	prev Hash // will be passed at runtime

	dataSum Hash // will be calculated at runtime

	created     [8]byte // uint64
	createdMsec [8]byte // uint64

	offset uint64
	length uint64
}

// EncodeBytes.
//
// [1] [2] [40] [8] [8] [8] [8]
//  ^   ^    ^   ^   ^   ^   ^
//  |   |    |   |   |   |   |
//  |   |    |   |   |   |   --- data length
//  |   |    |   |   |   ------- data offset pointer
//  |   |    |   |   ----------- creation timestamp msec
//  |   |    |   --------------- creation timestamp sec
//  |   |    ------------------- block hash
//  |   ------------------------ header length (without first 3 bytes)
//  ---------------------------- version

func (b *block) EncodeBytes(into []byte) error {

	return nil
}

func (b *block) DecodeBytes(from []byte) error {

	return nil
}

type data struct {
	mm mmap.MMap
}
