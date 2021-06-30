package consumer

import (
	"io"
	"os"
)

const (
	BufferSize = 4 * 1024
	PacketSize = 4 * 1024 * 1024
)

var ACK = []byte{0}

type Consumer interface {
	Handle(io.ReadWriteCloser, *os.File) error
}

type consumer struct {
}

func New() Consumer {
	return &consumer{}
}

func (c *consumer) Handle(rwc io.ReadWriteCloser, wal *os.File) error {
	defer rwc.Close()

	ack := make(chan struct{}, 100)
	done := make(chan struct{}, 1)

	defer close(done)

	go func() {
		for {
			select {
			case <-ack:
				rwc.Write(ACK)
			case <-done:
				return

			}
		}
	}()

	buf := make([]byte, BufferSize)
	total := 0
	for {
		rdr := io.TeeReader(rwc, wal)
		n, err := rdr.Read(buf)
		if n == 0 || err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}
		total += len(buf[:n])

		if total >= PacketSize {
			wal.Sync()
			total -= PacketSize
			ack <- struct{}{}
		}
	}
}
