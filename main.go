package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
	"unicode"
)

type Typer struct {
	buf  *bytes.Buffer
	rate time.Duration
}

func NewTyper(rate time.Duration) Typer {
	return Typer{new(bytes.Buffer), rate}
}

func (t Typer) Write(p []byte) (n int, err error) {
	return t.buf.Write(p)
}

func (t Typer) Read(p []byte) (n int, err error) {
	next, err := t.buf.ReadByte()
	delay := t.rate
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
	typer := NewTyper(time.Millisecond * 100)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Fprintln(typer, scanner.Text())
		io.Copy(os.Stdout, typer)
	}
}
