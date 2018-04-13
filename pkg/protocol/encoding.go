package protocol

import (
	"encoding/binary"
	"io"
)

var Encoding = binary.BigEndian

func ReadUInt8A(r io.Reader, n *uint8) {
	binary.Read(r, Encoding, n)
}

func ReadUInt16A(r io.Reader, n *uint16) {
	binary.Read(r, Encoding, n)
}

func ReadUInt32A(r io.Reader, n *uint32) {
	binary.Read(r, Encoding, n)
}

func ReadUInt64A(r io.Reader, n *uint64) {
	binary.Read(r, Encoding, n)
}

func ReadUInt8B(r io.Reader, n *uint8) {
	var header = make([]byte, 1)

	_, err := io.ReadFull(r, header)
	if err != nil {
		panic(err)
	}

	*n = uint8(header[0])
}

func ReadUInt16B(r io.Reader, n *uint16) {
	var header = make([]byte, 2)

	_, err := io.ReadFull(r, header)
	if err != nil {
		panic(err)
	}

	*n = Encoding.Uint16(header[:])
}

func ReadUInt32B(r io.Reader, n *uint32) {
	var header = make([]byte, 4)

	_, err := io.ReadFull(r, header)
	if err != nil {
		panic(err)
	}

	*n = Encoding.Uint32(header[:])
}

func ReadUInt64B(r io.Reader, n *uint64) {
	var header = make([]byte, 8)

	_, err := io.ReadFull(r, header)
	if err != nil {
		panic(err)
	}

	*n = Encoding.Uint64(header[:])
}
