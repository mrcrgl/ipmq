# Topics

- [Product spec](#product-spec)
  - what to expect?
- Protocol design
  - Headers, ...
- Command & Flow Control (when and where to forward etc.)
  - Stream data uninterrupted to the next segment if one is exchaused?
- Storage design
  - Commitchain
  - Segments & Shards
  - Blockchain and its advantages
- CommitLog Reference design and segemnt recognition
  - How to route to the next segment?
  - How to validate the data?
- Maybe how the components stick together?
- Component graph
- APIs
  - Consume
  - Produce

## Product Spec

1. [eventually consistent](https://en.wikipedia.org/wiki/Eventual_consistency)
2. Strict order
3. Consumer Groups
4. Read replicas (consumer devices behind firewall)
5. Highly distributed & operational cloud proven (robust design)

## Storage Design

Each piece of data is called a **[block](#block)** and is identified by it's **[id](#block-id)**. Blocks are chained with it's previous one (Blockchain).

**[Blocks](#block)** are chained together within a **[Segment](#segment)**. Either a number or a byte limit described in the **[Segment](#segment)** limits it's capacity. If a **[Block](#block)** doesn't fit in the current **[Segment](#segment)**, a new one will be created.

**[Segments](#segment)** itself are organized in **[Shards](#shard)**, 1 writable and 0..n readable. Every reading **[Shard](#shard)** may be elected to the leader (writable), if the current leader is going to or is gone out of service. **[Shards](#shard)** should be distributed across AZ's. **[Segments](#segment)** are chained together with it's previous one.

The external reference of a Blockchain is called **[CommitChain](#commitchain)**. It holds the hash to identify the first **[Segment](#segment)**.

### Block

Blocks are made of two parts. The header and the data. Both are hold by the related **Shard**.

#### Header v1

Total header bytes 101

| 4 bytes | 1 byte | 32 bytes | 32 bytes | 16 bytes | 8 bytes | 8 bytes |
|---|---|---|---|---|---|---|
| Header length - uint32 | Version - uint8 | Hash - sha256 | Prev. Hash - sha256 | Timestamp with nanosec - 2x uint64 | Data offset - uint64 | Data length - uint64 |

### Block ID

The ID consists of 2 parts, joined by a dot. The first part is a slice of the first 8 bytes of the **Segments Hash** the Block belongs to. The second one hashes the block itself.

<8 byte segment hash>.<32 byte block hash>

#### Block is hashed like this

SHA256: <32 byte prev block hash><16 byte timestamp><0..n byte data>

### Shard

??? Shard manage the read-write operations to handle with blocks.
It manages two byte arrays, one to the block index and one for the data of the blocks. The ladder is a memory map for faster lookups.

Properties
- Segment Hash
- ReadWrite/Read (!!! This should tell someone else)
- Shard No

### Segment

Segments are logical buckets. They provide an interface for reading and writing Blocks. Every request, reading or writing will be
- handled locally if the correct shard is available, or (???) <- this requires shared state, better to route before?
- proxied to the next host

Segments are identified by it's id, an 8 byte string made of ???

!TODO!
The first segment has the same timestamp as the CommitChain and therefore the same Hash?
What about evicted Segments?
Do Segments last forever and just Shards get evicted?
How to proof consistency? Should we hash the containing data?
What's the first Block hash made of?

Properties
- CommitChain Hash
- Hash
- Prev. Segment Hash
- Timestamp

### CommitChain

This describes the topic, the data itself. It has a local id that consists of an 8 byte string and a named string with a maximum of 128 bytes.

Properties
- Hash
- String identifier
- Timestamp

# Structure

## pkg/api/...

IPMQ api implementation, you can query for a certain Kind

```go

// Handler maybe of ReadCloser & WriteCloser?
handler := api.Scheme.Lookup(kind, version)
handler.Handle(...)

```

## pkg/apis/...

