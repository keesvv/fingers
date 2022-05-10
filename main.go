package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

type Typer struct {
	buf *bytes.Buffer
}

func NewTyper() Typer {
	return Typer{new(bytes.Buffer)}
}

func (t Typer) Write(p []byte) (n int, err error) {
	return t.buf.Write(p)
}

func (t Typer) Read(p []byte) (n int, err error) {
	time.Sleep(time.Millisecond * 100)
	next, err := t.buf.ReadByte()
	return copy(p, []byte{next}), err
}

func main() {
	typer := NewTyper()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Fprintln(typer, scanner.Text())
		io.Copy(os.Stdout, typer)
	}
}
