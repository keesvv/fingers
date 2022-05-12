package keyboard

type Layout struct {
	Keys [][]rune
}

var QwertyLayout = &Layout{[][]rune{
	[]rune("qwertyuiop"),
	[]rune("asdfghjkl"),
	[]rune("zxcvbnm"),
}}

func GetLayoutByID(id string) *Layout {
	if id == "qwerty" {
		return QwertyLayout // FIXME
	}
	return nil
}

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

func (l *Layout) GetKey(x, y int) rune {
	return l.Keys[y][x]
}

func (l *Layout) GetAdjacent(r rune, x, y int) rune {
	oX, oY := l.GetLoc(r)
	return l.GetKey(oX+x, oY+y)
}
