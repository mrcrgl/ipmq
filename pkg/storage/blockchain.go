package storage

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func NewBlockchain() *Blockchain {
	return &Blockchain{
		blocks: []*Block{
			NewBlock([]byte(""), []byte{}),
		},
	}
}

// Blockchain
type Blockchain struct {
	SegmentSize uint64
	segments    []*Segment
	blocks      []*Block // <- this may be a bad idea, move it to segment?
}

func (bc *Blockchain) Add(data []byte) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	// TODO This is kinda unperformant
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) Hashes() [][]byte {
	out := make([][]byte, len(bc.blocks))
	for i, b := range bc.blocks {
		out[i] = b.Hash[:]
	}
	return out
}

func (bc *Blockchain) newSegment(startHash []byte) error {
	s, err := NewSegment(startHash)
	if err != nil {
		return err
	}

	bc.segments = append(bc.segments, s)

	return nil
}

func NewSegment(startHash []byte) (*Segment, error) {
	dataStor, err := os.OpenFile(fmt.Sprintf("logPath/%x.blob", startHash), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	s := &Segment{
		StartHash: startHash,
		data:      dataStor,
		reader:    dataStor,
		writer:    dataStor,
		maxBytes:  10 * (1024 ^ 2),
	}

	return s, nil
}

// Segment is some slice of data within a Channel. The difficulty is to route to the
// correct Segment - how does a node know where to find the data or where to save?
type Segment struct {
	StartHash []byte
	reader    io.Reader
	writer    io.Writer
	data      *os.File
	blocks    []*Block
	maxBytes  int64
}

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

type Shard struct{}

func NewBlock(data, hash []byte) *Block {
	b := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          data,
		PrevBlockHash: hash,
	}

	b.SetHash()

	return b
}
