package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
	"time"
	"unicode"
)

type Typer struct {
	buf       *bytes.Buffer
	bps       uint8
	precision uint8
	last      byte
}

func NewTyper(bps, precision uint8) *Typer {
	return &Typer{new(bytes.Buffer), bps, precision, 0}
}

func (t *Typer) typo() bool {
	if t.precision >= 100 {
		return false
	}

	if t.precision == 0 {
		return true
	}

	n, err := rand.Int(rand.Reader, big.NewInt(99))
	if err != nil {
		panic(err)
	}

	cmp := n.Cmp(big.NewInt(int64(t.precision)))
	return cmp == 1 || cmp == 0
}

func (t *Typer) Write(p []byte) (n int, err error) {
	for _, b := range p {
		t.buf.WriteByte(b)
		n++

		if unicode.IsLetter(rune(b)) && t.typo() {
			typo := []byte{'\b', b}
			t.buf.Write(typo)
			n += len(typo)
		}
	}
	return
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
	} else if rLast == '\b' {
		delay *= 3
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
	precision := flag.Uint("p", 90, "precision percentage")
	flag.Parse()

	typer := NewTyper(uint8(*rate), uint8(*precision))
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Fprintln(typer, scanner.Text())
		io.Copy(os.Stdout, typer)
	}
}
