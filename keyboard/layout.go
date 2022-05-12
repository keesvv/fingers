package keyboard

type Layout [][]rune

var QwertyLayout = &Layout{
	[]rune("qwertyuiop"),
	[]rune("asdfghjkl"),
	[]rune("zxcvbnm"),
}

func GetLayoutByID(id string) *Layout {
	if id == "qwerty" {
		return QwertyLayout // FIXME
	}
	return nil
}

func (l *Layout) GetLoc(r rune) (x, y int) {
	for rowIndex, row := range *l {
		for keyIndex, key := range row {
			if r == key {
				return keyIndex, rowIndex
			}
		}
	}
	return -1, -1
}

func (l *Layout) GetKey(x, y int) rune {
	if x < 0 || y < 0 || y > len(*l)-1 || x > len((*l)[y])-1 {
		return 0
	}
	return (*l)[y][x]
}

func (l *Layout) GetAdjacent(r rune, x, y int) rune {
	oX, oY := l.GetLoc(r)
	return l.GetKey(oX+x, oY+y)
}
