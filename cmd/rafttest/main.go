package main

import (
	"flag"
	"net"
	"os"
	"strings"
	"time"

	"sync"

	"fmt"

	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-mdb"
)

var _ raftmdb.MDBStore

var (
	peers  *string
	listen *string
)

func init() {
	peers = flag.String("peers", "", "address:port,address2:port2")
	listen = flag.String("listen", ":8001", "address:port")

	flag.Parse()
}

func main() {
	/*mdb, err := raftmdb.NewMDBStore("test-store.mdb")
	if err != nil {
		panic(err)
	}*/

	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(*listen)

	peerList := strings.Split(*peers, ",")
	servers := make([]raft.Server, len(peerList))
	for i, peer := range peerList {
		servers[i] = raft.Server{
			raft.Voter,
			raft.ServerID(peer),
			raft.ServerAddress(peer),
		}
	}

	configuration := raft.Configuration{Servers: servers}

	addr, err := net.ResolveTCPAddr("tcp", *listen)
	if err != nil {
		panic(err)
	}

	/*l, err := net.ListenTCP("ipv4", addr)
	if err != nil {
		panic(err)
	}*/

	transport, err := raft.NewTCPTransport(addr.String(), addr, 3, time.Second*1, os.Stdout)
	if err != nil {
		panic(err)
	}

	/*err = raft.BootstrapCluster(
		config,
		raft.NewInmemStore(),
		raft.NewInmemStore(),
		raft.NewDiscardSnapshotStore(),
		transport,
		configuration,
	)
	if err != nil {
		panic(err)
	}*/

	//fsm := fuzzy.LoggerAdapter{}

	r, err := raft.NewRaft(
		config,
		nil,
		raft.NewInmemStore(),
		raft.NewInmemStore(),
		raft.NewDiscardSnapshotStore(),
		transport,
	)
	if err != nil {
		panic(err)
	}

	f := r.BootstrapCluster(configuration)
	if f.Error() != nil {
		panic(f.Error())
	}

	oCh := make(chan raft.Observation)

	or := raft.NewObserver(oCh, true, func(or *raft.Observation) bool {
		switch or.Data.(type) {
		case raft.RaftState:
			return true
		default:
			return false
		}
	})

	r.RegisterObserver(or)

	// r.Apply()

	go func() {
		for {
			rs := <-oCh
			if state, ok := rs.Data.(raft.RaftState); ok {
				fmt.Println("---> Got state: ", state.String())
			}
		}
	}()

	go func() {
		for {
			f := r.VerifyLeader()
			e := f.Error()
			switch e {
			case nil:
				fmt.Println(time.Now().Format(time.RFC3339Nano), " >> i am legend")
			case raft.ErrNotLeader:
				fmt.Println(time.Now().Format(time.RFC3339Nano), " >> '-(")
			default:
				fmt.Println(time.Now().Format(time.RFC3339Nano), "wtf is: ", e)
			}

			time.Sleep(time.Second * 2)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
