package commitchain

// Topic is a collection of segments
type Topic struct {
	// name of the stream (unique)
	name string

	// hash of the stream, represents the internal identifier
	hash []byte // how many bytes?

	// writable indicates whether the stream is writable or not
	writable bool

	// shards sets the required amount of replicas
	shards int // is this setting duplicated over all segments?

	// segments of the stream,
	segments []segment // necessary? Order need to be synced across nodes which can be tricky

	// evictionPolicy defines the eviction policy of bound segments
	evictionPolicy int // TODO change type
}
