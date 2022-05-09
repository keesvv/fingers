package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

type Typer struct {
	buf *bytes.Buffer
}

func NewTyper() Typer {
	return Typer{bytes.NewBuffer(make([]byte, 0))}
}

func (t Typer) Write(p []byte) (n int, err error) {
	return t.buf.Write(p)
}

func (t Typer) Read(p []byte) (n int, err error) {
	time.Sleep(time.Millisecond * 100)
	next, err := t.buf.ReadByte()

	if err != nil {
		return 0, err
	}

	copy(p, []byte{next})
	return 1, nil
}

func main() {
	typer := NewTyper()
	io.Copy(typer, os.Stdin)
	io.Copy(os.Stdout, typer)
}
