package draw

type Font struct {
	Height int

	// It is allowed to have a smaller line height than height.
	LineHeight int

	charWidths map[rune]float32
}

// TextWidth returns the width of the given text based on a 12px arial
// font.
func (f Font) TextWidth(txt string) int {
	var width float32
	for _, r := range txt {
		w, found := f.charWidths[r]
		if !found {
			w = 8.0
		}
		width += float32(w) * float32(f.Height) / float32(12)
	}
	return int(width)
}

// font-size: 12px, most ascii characters
var arial = map[rune]float32{
	'!':  3,
	'"':  4,
	'#':  7,
	'$':  7,
	'%':  11,
	'&':  8,
	'\'': 2,
	'(':  4,
	')':  4,
	'*':  5,
	'+':  7,
	',':  3,
	'-':  4,
	'.':  3,
	'/':  3,
	'0':  7,
	'1':  7,
	'2':  7,
	'3':  7,
	'4':  7,
	'5':  7,
	'6':  7,
	'7':  7,
	'8':  7,
	'9':  7,
	':':  3,
	';':  3,
	'<':  7,
	'=':  7,
	'>':  7,
	'?':  7,
	'@':  12,
	'A':  8,
	'B':  8,
	'C':  9,
	'D':  9,
	'E':  8,
	'F':  7,
	'G':  9,
	'H':  9,
	'I':  3,
	'J':  6,
	'K':  8,
	'L':  7,
	'M':  10,
	'N':  9,
	'O':  9,
	'P':  8,
	'Q':  9,
	'R':  9,
	'S':  8,
	'T':  7,
	'U':  9,
	'V':  8,
	'W':  11,
	'X':  8,
	'Y':  8,
	'Z':  7,
	'[':  3,
	'\\': 3,
	']':  3,
	'^':  5,
	'_':  7,
	'`':  4,
	'a':  7,
	'b':  7,
	'c':  6,
	'd':  7,
	'e':  7,
	'f':  3,
	'g':  7,
	'h':  7,
	'i':  3,
	'j':  3,
	'k':  6,
	'l':  3,
	'm':  10,
	'n':  7,
	'o':  7,
	'p':  7,
	'q':  7,
	'r':  4,
	's':  6,
	't':  3,
	'u':  7,
	'v':  6,
	'w':  9,
	'x':  6,
	'y':  6,
	'z':  6,
	'{':  4,
	'|':  3,
	'}':  4,
	'~':  7,
	' ':  3,
	'¡':  4,
	'¢':  7,
	'£':  7,
	'¤':  7,
	'¥':  7,
	'¦':  3,
	'§':  7,
	'¨':  4,
	'©':  9,
	'ª':  5,
	'«':  7,
	'¬':  7,
	'­':  0,
	'®':  9,
	'¯':  7,
	'°':  5,
	'±':  7,
	'²':  4,
	'³':  4,
	'´':  4,
	'µ':  7,
	'¶':  7,
	'·':  4,
	'¸':  4,
	'¹':  4,
	'º':  4,
	'»':  7,
	'¼':  10,
	'½':  10,
	'¾':  10,
	'¿':  7,
	'À':  8,
	'Á':  8,
	'Â':  8,
	'Ã':  8,
	'Ä':  8,
	'Å':  8,
	'Æ':  12,
	'Ç':  9,
	'È':  8,
	'É':  8,
	'Ê':  8,
	'Ë':  8,
	'Ì':  3,
	'Í':  3,
	'Î':  3,
	'Ï':  3,
	'Ð':  9,
	'Ñ':  9,
	'Ò':  9,
	'Ó':  9,
	'Ô':  9,
	'Õ':  9,
	'Ö':  9,
	'×':  7,
	'Ø':  9,
	'Ù':  9,
	'Ú':  9,
	'Û':  9,
	'Ü':  9,
	'Ý':  8,
	'Þ':  8,
	'ß':  7,
	'à':  7,
	'á':  7,
	'â':  7,
	'ã':  7,
	'ä':  7,
	'å':  7,
	'æ':  11,
	'ç':  6,
	'è':  7,
	'é':  7,
	'ê':  7,
	'ë':  7,
	'ì':  3,
	'í':  3,
	'î':  3,
	'ï':  3,
	'ð':  7,
	'ñ':  7,
	'ò':  7,
	'ó':  7,
	'ô':  7,
	'õ':  7,
	'ö':  7,
	'÷':  7,
	'ø':  7,
	'ù':  7,
	'ú':  7,
	'û':  7,
	'ü':  7,
	'ý':  6,
	'þ':  7,
}
