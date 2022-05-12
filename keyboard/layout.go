package keyboard

type Layout struct {
	Keys [][]rune
}

var QwertyLayout = &Layout{[][]rune{
	[]rune("qwertyuiop"),
	[]rune("asdfghjkl"),
	[]rune("zxcvbnm"),
}}

func (l *Layout) GetLoc(r rune) (x, y int) {
	for rowIndex, row := range l.Keys {
		for keyIndex, key := range row {
			if r == key {
				return keyIndex, rowIndex
			}
		}
	}
	return -1, -1
}
