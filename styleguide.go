package design

import (
	"fmt"

	"github.com/gregoryv/go-design/xml"
)

type StyleGuide struct {
	applicable bool

	FontWidth     int
	LineHeight    int
	PaddingTop    int
	PaddingBottom int
	PaddingLeft   int
	Space         int // between components in graph

	ShapeStrokeWidth int
	ShapeStrokeColor string
	ShapeFill        string
}

var DefaultStyle = StyleGuide{
	applicable: true,

	FontWidth:     16,
	LineHeight:    12,
	PaddingTop:    4,
	PaddingBottom: 4,
	PaddingLeft:   16,
	Space:         60,

	ShapeStrokeWidth: 1,
	ShapeStrokeColor: "black",
	ShapeFill:        "#ffffcc",
}

func (s *StyleGuide) HasSpecialStyle() bool {
	return s.applicable
}

func (s *StyleGuide) FillStroke() xml.Attribute {
	return style(fmt.Sprintf("fill:%s;stroke:%s;stroke-width:%v",
		s.ShapeFill,
		s.ShapeStrokeColor,
		s.ShapeStrokeWidth,
	))
}

func (s *StyleGuide) Stroke() xml.Attribute {
	return style(fmt.Sprintf("stroke:%s;stroke-width:%v",
		s.ShapeStrokeColor,
		s.ShapeStrokeWidth,
	))
}

func widthOf(txt string) int {
	return len(txt) * DefaultStyle.FontWidth
}

func (s *StyleGuide) Height(lines int) int {
	h := 2 // below line, e.g. letter 'g'
	return lines * (h + s.LineHeight + s.PaddingTop + s.PaddingBottom)
}

func (s *StyleGuide) Offset(x, y int) *Offset {
	return &Offset{x, y, s}
}

type Offset struct {
	x, y int
	s    *StyleGuide
}

// Returns the y offset for line number n
func (o *Offset) Line(n int) int {
	return o.y + (o.s.LineHeight+o.s.PaddingTop+o.s.PaddingBottom)*n
}
