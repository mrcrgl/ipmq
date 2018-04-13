package server

import (
	"io"
	"log"
	"net"

	"bytes"

	"crypto/sha256"

	"github.com/mrcrgl/ipmq/src/protocol"
)

type handler struct {
	listener net.Listener
	network  string
	addr     string
	cCh      chan struct{}
	log      *log.Logger
}

func New(logger *log.Logger, network, addr string) *handler {
	h := new(handler)

	h.network = network
	h.addr = addr
	h.cCh = make(chan struct{})
	h.log = logger

	return h
}

func (h *handler) Listen() error {
	if err := h.setupListener(); err != nil {
		return err
	}

	return h.listen()
}

func (h *handler) Stop() error {
	close(h.cCh)

	// TODO wait for connections

	return nil
}

func (h *handler) setupListener() error {
	var err error
	if h.listener, err = net.Listen(h.network, h.addr); err != nil {
		return err
	}

	return nil
}

func (h *handler) listen() error {

	for {
		select {
		case <-h.cCh:
			break
		default:
			conn, err := h.listener.Accept()
			if err != nil {
				h.log.Printf("Failed to establish connection: %v", err)
			}

			go h.serve(conn)

		}
	}

	return nil
}

func (h *handler) serve(conn net.Conn) {
	addr := conn.RemoteAddr()
	h.log.Printf("peer %s connected", addr.String())

	for {
		/*if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
			if err == io.EOF {
				h.log.Printf("connection closed: %v", addr)
				break
			}

			h.log.Printf("read deadline failed: %s", err)
			continue
		}*/

		header := make([]byte, 6)
		if _, err := io.ReadFull(conn, header[:]); err != nil {
			if err == io.EOF {
				h.log.Printf("connection closed: %v", addr)
				break
			}

			h.log.Printf("failed to read header(apiKind): %v", err)
			break
		}

		//h.log.Printf("Header: %v", header)

		size := int64(protocol.Encoding.Uint32(header[0:4]))

		/*if size != 100 {
			h.log.Printf("Size mismatch: %d", size)
			break
		}*/

		apiKind := protocol.Encoding.Uint16(header[4:6])

		if apiKind != 8 {
			h.log.Printf("Invalid apiKind: %d", apiKind)
			break
		}

		handleProduce(io.LimitReader(conn, size), conn)

		hash := sha256.New()
		buf := bytes.NewBuffer(make([]byte, size))

		wr := io.MultiWriter(buf, hash)

		if n, err := io.CopyN(wr, conn, size); err != nil {
			if err == io.EOF {
				h.log.Printf("connection closed: %v", addr)
				break
			}

			h.log.Printf("Failed to read body: %v (%d bytes read)", err, n)
			break
		} else if n != size {
			h.log.Printf("Body size mismatch: expected %d but got %d", size, n)
			break
		}

		hash.Sum(nil)
		//h.log.Printf("%x", hash.Sum(nil))

		// lookup requested topic
		// get last offset
		// check remaining capacity
		// write bytes to mmap
		// when succeed, write index
		// respond to client

		//h.log.Printf("This is apiKind: %d", apiKind)

	}

	conn.Close()
}

type Request struct {
}

func handleProduce(r io.Reader, w io.Writer) {

}

func handleDebug(r io.Reader, w io.Writer) {
	/*hash := sha256.New()
	buf := bytes.NewBuffer(make([]byte, size))

	wr := io.MultiWriter(buf, hash)

	if n, err := io.CopyN(wr, conn, size); err != nil {
		if err == io.EOF {
			h.log.Printf("connection closed: %v", addr)
			break
		}

		h.log.Printf("Failed to read body: %v (%d bytes read)", err, n)
		break
	} else if n != size {
		h.log.Printf("Body size mismatch: expected %d but got %d", size, n)
		break
	}

	hash.Sum(nil)*/
}
