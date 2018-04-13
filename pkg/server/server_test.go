package server

import (
	"bytes"
	"io"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"fmt"

	"github.com/mrcrgl/ipmq/src/protocol"
)

func BenchmarkServe(b *testing.B) {
	clientData := make([]byte, 4+2+100)

	fmt.Printf("clientData length: %d", len(clientData))

	protocol.Encoding.PutUint32(clientData[0:4], 100)
	protocol.Encoding.PutUint16(clientData[4:6], 8)

	copy(clientData[6:], []byte("Hallo, this should be at least 100 bytes. Sometimes you win, sometimes not. That's fuckin´ life ^_^"))

	stdout := []byte{}
	var w io.Writer = os.Stderr // bytes.NewBuffer(stdout)
	h := New(log.New(w, "test: ", log.LstdFlags), "tcp", ":12312")

	/*b.Run("Server Put (with connect)", func(b *testing.B) {
		conn := &testConn{
			data: bytes.NewBuffer(clientData),
		}

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			h.serve(conn)
		}

		b.StopTimer()

		fmt.Println(stdout)
	})*/

	b.Run("Server Put (single connection)", func(b *testing.B) {
		buf := bytes.NewBuffer([]byte{})
		conn := &testConn{
			data: buf,
		}

		h.serve(conn)

		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			io.Copy(buf, bytes.NewBuffer(clientData))
		}

		b.StopTimer()

		fmt.Println(stdout)
	})

}

var _ net.Conn = &testConn{}

type testConn struct {
	data io.Reader
}

func (c *testConn) Close() error                       { return nil }
func (c *testConn) Read(b []byte) (int, error)         { return c.data.Read(b) }
func (c *testConn) Write(b []byte) (int, error)        { return 0, nil }
func (c *testConn) LocalAddr() net.Addr                { return &addr{"tcp", "0.0.0.0:12345"} }
func (c *testConn) RemoteAddr() net.Addr               { return &addr{"tcp", "0.0.0.0:12300"} }
func (c *testConn) SetDeadline(t time.Time) error      { return nil }
func (c *testConn) SetReadDeadline(t time.Time) error  { return nil }
func (c *testConn) SetWriteDeadline(t time.Time) error { return nil }

var _ net.Addr = &addr{}

type addr struct {
	network string
	address string
}

func (a *addr) Network() string {
	return a.network
}

func (a *addr) String() string {
	return a.address
}
