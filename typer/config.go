package typer

import "keesvv.nl/fingers/keyboard"

type Config struct {
	Layout  *keyboard.Layout
	Autofix bool
	Bps,
	Precision uint8
}
