package protocol

type ByteEncoder interface {
	ByteEncode()
}

type ByteDecoder interface {
	ByteDecode()
}
