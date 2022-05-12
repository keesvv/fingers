package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"keesvv.nl/fingers/typer"
)

func main() {
	rate := flag.Uint("r", 10, "average amount of bps (bytes per second)")
	precision := flag.Uint("p", 90, "precision percentage")
	flag.Parse()

	t := typer.NewTyper(uint8(*rate), uint8(*precision))
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Fprintln(t, scanner.Text())
		io.Copy(os.Stdout, t)
	}
}
