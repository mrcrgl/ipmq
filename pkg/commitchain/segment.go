package commitchain

type segment struct {
	hash Hash

	// TODO necessary?
	topic Hash

	// closed segments doesn't accept further data, another segment continues the stream
	closed bool

	shards []shard
}
