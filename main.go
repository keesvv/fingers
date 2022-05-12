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
	autofix := flag.Bool("f", true, "autofix typos")
	flag.Parse()

	kbLayout := keyboard.GetLayoutByID(*layout)
	if kbLayout == nil {
		panic(errors.New("no such layout")) // FIXME
	}

	t := typer.NewTyper(&typer.Config{
		Layout:    kbLayout,
		Autofix:   *autofix,
		Bps:       uint8(*rate),
		Precision: uint8(*precision),
	})

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintln(t, scanner.Text())
		io.Copy(os.Stdout, t)
	}
}
