package typer

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"time"
	"unicode"
)

type Typer struct {
	buf    *bytes.Buffer
	config *Config
	last   byte
}

func NewTyper(config *Config) *Typer {
	return &Typer{new(bytes.Buffer), config, 0}
}

func (t *Typer) typo() bool {
	if t.config.Precision >= 100 {
		return false
	}

	if t.config.Precision == 0 {
		return true
	}

	n, err := rand.Int(rand.Reader, big.NewInt(99))
	if err != nil {
		panic(err)
	}

	return n.Cmp(big.NewInt(int64(t.config.Precision))) >= 0
}

func (t *Typer) Write(p []byte) (n int, err error) {
	var buf []byte

	for _, b := range p {
		if unicode.In(
			rune(b), unicode.Letter, /*unicode.Digit,*/
		) && t.typo() {
			typo := t.config.Layout.GetAdjacent(rune(b), -1, 0)
			if typo == 0 {
				typo = rune(b)
			}

			buf = append(buf, []byte{byte(typo), '\b', b}...)
			continue
		}

		buf = append(buf, b)
	}

	return t.buf.Write(buf)
}

func (t *Typer) Read(p []byte) (n int, err error) {
	next, err := t.buf.ReadByte()
	if err != nil {
		return 0, err
	}

	rNext, rLast := rune(next), rune(t.last)
	delay := time.Second / time.Duration(t.config.Bps)
	switch {
	case rNext == rLast && unicode.IsGraphic(rNext), unicode.IsSpace(rNext):
		delay /= 2
	case unicode.IsDigit(rNext):
		delay *= 2
	case rNext == '\b':
		delay *= 3
	case unicode.IsGraphic(rNext):
		break
	default:
		delay = 0
	}

	t.last = next
	time.Sleep(delay)
	return copy(p, []byte{next}), nil
}
