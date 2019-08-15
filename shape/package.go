package shape

import (
	"bytes"
	"text/template"
)

type Stringer interface {
	String() string
}

type svg interface {
	Svg() string
}

func toString(xml string, shape interface{}) string {
	svg := template.Must(template.New("").Parse(xml))
	buf := bytes.NewBufferString("")
	svg.Execute(buf, shape)
	return buf.String()
}
