package circular

import "errors"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

type Buffer struct {
	data     []byte
	capacity int
	size     int
	next     int
	first    int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		data:     make([]byte, size),
		capacity: size,
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.size == 0 {
		return 0, errors.New("no stuff")
	}
	val := b.data[b.first]
	b.data[b.first] = 0
	b.size--
	b.first++
	if b.first >= b.capacity {
		b.first = 0
	}
	return val, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.size == b.capacity {
		return errors.New("too much stuff")
	}
	b.data[b.next] = c
	b.size++
	b.next++
	if b.next >= b.capacity {
		b.next = 0
	}
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	b.data[b.next] = c
	if b.size < b.capacity {
		b.size++
	} else {
		b.first++
		if b.first >= b.capacity {
			b.first = 0
		}
	}
	b.next++
	if b.next >= b.capacity {
		b.next = 0
	}
}

func (b *Buffer) Reset() {
	b.size = 0
	b.next = 0
	b.first = 0
	for i := range b.data {
		b.data[i] = 0
	}
}
