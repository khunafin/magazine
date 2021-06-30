package producer

import (
	"io"

	"github.com/khunafin/magazine/internal/list"
)

const PacketSize = 4 * 1024 * 1024

type Producer interface {
	Produce() error
}

type producer struct {
	rw    io.ReadWriter
	count int

	ll list.List
}

func (p *producer) Produce() error {
	add := make(chan []byte, 1000)
	del := make(chan struct{}, 1000)

	done := make(chan struct{}, 1)
	defer close(done)

	go func() {
		for {
			select {
			case payload := <-add:
				p.ll.Append(payload)
			case <-del:
				p.ll.Shrink()
			case <-done:
				return
			}
		}
	}()

	go func() {
		bfr := make([]byte, 1024)
		for {
			n, err := p.rw.Read(bfr)
			if n == 0 || err != nil {
				break
			}
			for j := 0; j < n; j++ {
				del <- struct{}{}
			}
		}
	}()

	for i := 0; i < p.count; i++ {
		payload := make([]byte, PacketSize)
		p.rw.Write(payload)
		add <- payload
	}

	return nil
}

func New(rw io.ReadWriter, count int) Producer {
	return &producer{
		rw:    rw,
		count: count,
		ll:    list.List{},
	}
}
