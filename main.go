package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"keesvv.nl/fingers/keyboard"
	"keesvv.nl/fingers/typer"
)

func main() {
	rate := flag.Uint("r", 10, "average amount of bps (bytes per second)")
	precision := flag.Uint("p", 90, "precision percentage")
	layout := flag.String("l", "qwerty", "keyboard layout")
	flag.Parse()

	kbLayout := keyboard.GetLayoutByID(*layout)
	if kbLayout == nil {
		panic(errors.New("no such layout")) // FIXME
	}

	t := typer.NewTyper(uint8(*rate), uint8(*precision), kbLayout)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Fprintln(t, scanner.Text())
		io.Copy(os.Stdout, t)
	}
}
