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
	buf  *bytes.Buffer
	bps  uint8
	last byte
}

func NewTyper(bps uint8) *Typer {
	return &Typer{new(bytes.Buffer), bps, 0}
}

func (t *Typer) Write(p []byte) (n int, err error) {
	return t.buf.Write(p)
}

func (t *Typer) Read(p []byte) (n int, err error) {
	next, err := t.buf.ReadByte()
	if errors.Is(err, io.EOF) {
		return 0, err
	}

	rNext, rLast := rune(next), rune(t.last)
	delay := time.Second / time.Duration(t.bps)
	if rNext == rLast {
		delay /= 2
	} else if unicode.IsSpace(rNext) {
		delay /= 3
	} else if unicode.IsDigit(rNext) {
		delay *= 2
	}

	t.last = next
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
