package paasio

import (
	"io"
	"sync"
)

type (
	writeCounter struct {
		writer io.Writer
		bytes  int64
		ops    int
		mu     *sync.Mutex
	}

	readCounter struct {
		reader io.Reader
		bytes  int64
		ops    int
		mu     *sync.Mutex
	}

	readWriteCounter struct {
		readCounter  ReadCounter
		writeCounter WriteCounter
	}
)

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer, mu: &sync.Mutex{}}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader, mu: &sync.Mutex{}}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter:  NewReadCounter(readwriter),
		writeCounter: NewWriteCounter(readwriter),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	n, err := rc.reader.Read(p)
	rc.bytes += int64(n)
	rc.ops++
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.bytes, rc.ops
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	n, err := wc.writer.Write(p)
	wc.bytes += int64(n)
	wc.ops++
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return wc.bytes, wc.ops
}

func (rwc *readWriteCounter) Read(p []byte) (int, error) {
	return rwc.readCounter.Read(p)
}

func (rwc *readWriteCounter) Write(p []byte) (int, error) {
	return rwc.writeCounter.Write(p)
}

func (rwc *readWriteCounter) ReadCount() (int64, int) {
	return rwc.readCounter.ReadCount()
}

func (rwc *readWriteCounter) WriteCount() (int64, int) {
	return rwc.writeCounter.WriteCount()
}
