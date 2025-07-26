package iox

import (
	"io"
	"os"
	"sync/atomic"
)

type WriterFunc func(p []byte) (n int, err error)

func (w WriterFunc) Write(p []byte) (n int, err error) {
	return w(p)
}

func CheckingPipe(w io.Writer) (io.Writer, *atomic.Bool) {
	var result atomic.Bool

	return WriterFunc(func(p []byte) (n int, err error) {
		result.Store(true)
		return w.Write(p)
	}), &result
}

func Exists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
