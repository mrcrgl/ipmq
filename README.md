
* Blockchain
  * has an identifier
  * has many Segments
  * locked (no writes permitted, for fixed end blobs for instance)
  * TODO find better name

Segments of a Blockchain may be distributed over different devices, so everytime you write into a blockchain, 
the corresponding segment and it's master must be discovered and the traffic routed to.
=> Which node will create the new Segment? 
=> Do we have enough space for it?
=> Are the read replicas necessary?

* Segment
  * has many Blocks
  * has an index and a data file
  * has a fixed length or size (must be size afaik since the mmap requires it)

* Block
  * has many bytes of data, timestamp, prev-hash and it's own hash

## Goals

[ ] Resilient and cloud proven design
    [ ] Scaling capability
    [ ] Auto failover
    [ ] Built in cluster management
[ ] Accurate order of messages
[ ] High throughput & low latency
[ ] Flexible distribution of data
[ ] Easy validation of messages
[ ] Built in support of data streams

What can you expect
* Data is mostly safe in case of failure (shard replication)
* Write order is proven
* Write persistence is proven
* Write replication may be proven (depending on settings?)
* Message id lookup may take some time
* Reading streams, even concurrent, is blazing fast

## Features

[ ] 

## Terminology

Stream
Pipeline

+------------------------------------------------------+
|                                                      |
|                        +---------------------------+ |
|                        | Segment [hash 1]          | |
|                        |     - Block [hash 2]      | |
|                        |     - Block [hash 3]      | |
|                        |     - Block [hash 4]      | |
|                        +---------------------------+ |
|                                                      |
|                        +---------------------------+ |
|                        | Segment [hash 4]          | |
|                        |     - Block [hash 5]      | |
|                        |     - Block [hash 6]      | |
|                        |     - Block [hash 7]      | |
|                        +---------------------------+ |
|                                                      |
|                        +---------------------------+ |
|                        | Segment [hash 7]          | |
|                        |     - Block [hash 8]      | |
|                        |     - Block [hash 9]      | |
|                        |     - Block [hash 10]     | |
|                        +---------------------------+ |
+------------------------------------------------------+

### Hashing proposal

segment hash
block hash

shash1.bhash1
shash1.bhash2
shash1.bhash3
shash1.bhash4

shash2 = bhash4
shash2.bhash1
shash2.bhash2
shash2.bhash3
shash2.bhash4

## Things to solve

* What is the name of the topmost struct (called r2d2 for now)?
* How is the information shared across the nodes...
  * which segments belongs to corresponding r2d2?
  * what's the last segment of given r2d2?
  * r2d2 settings...

## Contrast to Kafka

### Kafka provides
- Partitions? (Ordering is garanteed per partition in kafka)

- Consumer Groups deliver each message once per group

https://sookocheff.com/post/kafka/kafka-in-a-nutshell/

Kafka guarantees:
(1) Messages sent to a topic partition will be appended to the commit log in the order they are sent,
(2) a single consumer instance will see messages in the order they appear in the log,
(3) a message is ‘committed’ when all in sync replicas have applied it to their log, and
(4) any committed message will not be lost, as long as at least one in sync replica is alive

Kafka Produce Client may either
(1) wait for all in sync replicas to acknowledge the message,
(2) wait for only the leader to acknowledge the message, or
(3) do not wait for acknowledgement

Kafka Consume Client may
(1) receive each message at most once, 
(2) receive each message at least once, or 
(3) receive each message exactly once

Kafka Performance: https://engineering.linkedin.com/kafka/benchmarking-apache-kafka-2-million-writes-second-three-cheap-machines

## API

Stream      - Continuous read stream of data

Poll        - Ask for Blocks since last token, releases client after fetching remaining blocks
LongPolling - Get or Wait for Blocks since last token, releases client after fetching remaining blocks

Fetch       - Get a single Block by token
FetchBatch  - Get a range of Blocks between tokens


* Administrative
  * 
* Peering
  * 
* Consumer
  * 

# Architecture of the cluster

-> is it possible to skip the internal metadata storage?

A cluster may have 
- many streams, which have
- many segments, which have
- a certain amount of shards


How to bootstrap the cluster
* on kubernetes
  * address lack of stability
* local notebook
* existing etcd cluster

What information need to be stored
* For 500 streams with 100 Segments and approximately 3 shards, it's calculated to be around 11 Mib


sha1 == [20]byte


