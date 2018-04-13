package protocol

import (
	"testing"
)

type r struct {
	b []byte
}

var (
	sampleUInt8  = []byte{128}
	sampleUInt16 = []byte{128, 128}
	sampleUInt32 = []byte{128, 128, 128, 128}
	sampleUInt64 = []byte{128, 128, 128, 128, 128, 128, 128, 128}
)

func (r *r) Read(p []byte) (int, error) {
	copy(p, r.b)
	return len(r.b), nil
}

func BenchmarkReadUInt8A(b *testing.B) {

	var n uint8 = 0
	read := &r{sampleUInt8}

	for i := 0; i < b.N; i++ {
		ReadUInt8A(read, &n)
	}

	var _ = n
}

func BenchmarkReadUInt16A(b *testing.B) {

	var n uint16 = 0
	read := &r{sampleUInt16}

	for i := 0; i < b.N; i++ {
		ReadUInt16A(read, &n)
	}

	var _ uint16 = n
}

func BenchmarkReadUInt32A(b *testing.B) {

	var n uint32 = 0
	read := &r{sampleUInt32}

	for i := 0; i < b.N; i++ {
		ReadUInt32A(read, &n)
	}

	var _ = n
}

func BenchmarkReadUInt64A(b *testing.B) {

	var n uint64 = 0
	read := &r{sampleUInt64}

	for i := 0; i < b.N; i++ {
		ReadUInt64A(read, &n)
	}

	var _ = n
}

func BenchmarkReadUInt8B(b *testing.B) {

	var n uint8 = 0
	read := &r{sampleUInt8}

	for i := 0; i < b.N; i++ {
		ReadUInt8B(read, &n)
	}

	var _ = n
}

func BenchmarkReadUInt16B(b *testing.B) {

	var n uint16 = 0
	read := &r{sampleUInt16}

	for i := 0; i < b.N; i++ {
		ReadUInt16B(read, &n)
	}

	var _ = n
}

func BenchmarkReadUInt32B(b *testing.B) {

	var n uint32 = 0
	read := &r{sampleUInt32}

	for i := 0; i < b.N; i++ {
		ReadUInt32B(read, &n)
	}

	var _ = n
}

func BenchmarkReadUInt64B(b *testing.B) {

	var n uint64 = 0
	read := &r{sampleUInt64}

	for i := 0; i < b.N; i++ {
		ReadUInt64B(read, &n)
	}

	var _ = n
}
