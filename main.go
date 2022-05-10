package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
	"unicode"
)

type Typer struct {
	buf *bytes.Buffer
	bps uint8
}

func NewTyper(bps uint8) Typer {
	return Typer{new(bytes.Buffer), bps}
}

func (t Typer) Write(p []byte) (n int, err error) {
	return t.buf.Write(p)
}

func (t Typer) Read(p []byte) (n int, err error) {
	next, err := t.buf.ReadByte()
	delay := time.Second / time.Duration(t.bps)
	r := rune(next)

	if errors.Is(err, io.EOF) {
		return 0, err
	}

	if unicode.IsSpace(r) {
		delay /= 3
	} else if unicode.IsDigit(r) {
		delay *= 2
	}

	time.Sleep(delay)
	return copy(p, []byte{next}), err
}

func main() {
	rate := flag.Uint("r", 10, "average amount of bps (bytes per second)")
	flag.Parse()

	typer := NewTyper(uint8(*rate))
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Fprintln(typer, scanner.Text())
		io.Copy(os.Stdout, typer)
	}
}
