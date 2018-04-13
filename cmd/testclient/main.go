package main

import (
	"net"

	"log"

	"os"
	"os/signal"
	"time"

	"bytes"
	"io"

	"sync"

	"github.com/mrcrgl/ipmq/src/protocol"
)

func main() {
	//body := []byte("Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^")
	body := []byte("Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^" +
		"Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^" +
		"Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^" +
		"Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^" +
		"Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^" +
		"Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^" +
		"Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^" +
		"Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^" +
		"Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^" +
		"Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^")

	apiKind := make([]byte, 2)
	protocol.Encoding.PutUint16(apiKind, 8)

	size := make([]byte, 4)
	protocol.Encoding.PutUint32(size, uint32(len(body)))

	log.Printf("size bytes: %v", size)
	log.Printf("Foo is %d bytes long", len(body))

	conn, err := net.Dial("tcp", ":51892")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	var bw sync.RWMutex
	var msgsWritten uint64
	var bytesWritten float64
	started := time.Now()

	stats := func() {
		duration := time.Now().Sub(started)

		bw.RLock()
		bytesPerSecond := bytesWritten / duration.Seconds()
		msgsPerSecond := float64(msgsWritten) / duration.Seconds()
		bw.RUnlock()

		log.Printf(
			"Thats %.0f bytes/s\t%.2f kb/s\t%.2f mb/s\t%.3f gb/s\t%.2f messages/s",
			bytesPerSecond,
			bytesPerSecond/1024,
			bytesPerSecond/(1024*1024),
			bytesPerSecond/(1024*1024*1024),
			msgsPerSecond,
		)

		started = time.Now()
		bw.Lock()
		msgsWritten = 0
		bytesWritten = 0
		bw.Unlock()
	}

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, os.Kill)

		for {
			select {
			case <-time.After(time.Second * 10):
				stats()
			case <-c:
				stats()
				os.Exit(0)
			}
		}
	}()

	for {
		if n, err := conn.Write(size[:]); err != nil {
			log.Fatalf("Failed to write header(size): %v", err)
		} else {
			bw.Lock()
			bytesWritten += float64(n)
			bw.Unlock()
		}

		if n, err := conn.Write(apiKind[:]); err != nil {
			log.Fatalf("Failed to write header(apiKind): %v", err)
		} else {
			bw.Lock()
			bytesWritten += float64(n)
			bw.Unlock()
		}

		br := bytes.NewReader(body[:])

		if n, err := io.Copy(conn, br); err != nil {
			log.Fatalf("Failed to write body: %v", err)
		} else {
			bw.Lock()
			bytesWritten += float64(n)
			bw.Unlock()
		}

		bw.Lock()
		msgsWritten += 1
		bw.Unlock()
	}
}
