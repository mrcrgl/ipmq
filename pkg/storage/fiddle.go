package storage

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"

	mmap "github.com/edsrzf/mmap-go"
)

type Pool struct {
	closed   bool
	segments []*Segment
}

type Segment struct {
	id        []byte
	index     *Index
	data      *Data
	maxLength uint64
}

type Index struct {
}

type Data struct {
	mmap mmap.MMap
}

func (d *Data) Write(b []byte) (int, error) {
	return
}

func (d *Data) WriteAt(offset int, b []byte) (int, error) {
	return
}

type Item struct {
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
	DataOffset    uint64
	DataLength    uint64
}

func (i *Item) SetHash() {
	timestamp := []byte(strconv.FormatInt(i.Timestamp, 10))
	headers := bytes.Join([][]byte{i.PrevBlockHash, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	i.Hash = hash[:]
}

func NewItem(data, hash []byte) *Item {
	i := &Item{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: hash,
	}

	i.SetHash()

	return i
}
