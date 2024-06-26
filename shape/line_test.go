package shape

import (
	"fmt"
	"os"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/drewjya/draw"
)

func TestArrow_Height(t *testing.T) {
	// start and stop is the same, but head adds a height
	a := NewArrow(50, 50, 50, 50)
	if got := a.Height(); got == 0 {
		t.Error("unexpected height", got)
	}
}

func Test_arrowDirections(t *testing.T) {
	arrows := map[string]struct {
		*Line
		Direction
	}{
		"testdata/arrow_points_up_and_right.svg": {
			NewArrow(50, 50, 80, 20),
			DirectionUpRight,
		},
		"testdata/arrow_points_up_and_left.svg": {
			NewArrow(50, 50, 20, 20),
			DirectionUpLeft,
		},
		"testdata/arrow_points_down_and_left.svg": {
			NewArrow(50, 50, 40, 80),
			DirectionDownLeft,
		},
		"testdata/arrow_points_down_and_right.svg": {
			NewArrow(50, 50, 70, 80),
			DirectionDownRight,
		},
		"testdata/arrow_points_right.svg": {
			NewArrow(50, 50, 100, 50),
			DirectionRight,
		},
		"testdata/arrow_points_left.svg": {
			NewArrow(50, 50, 10, 50),
			DirectionLeft,
		},
		"testdata/arrow_points_down.svg": {
			NewArrow(50, 50, 50, 100),
			DirectionDown,
		},
		"testdata/arrow_points_up.svg": {
			NewArrow(50, 50, 50, 10),
			DirectionUp,
		},
	}

	for file, c := range arrows {
		t.Run(file, func(t *testing.T) {
			saveAs(t, c.Line, file)
			dir := c.Line.Direction()
			if dir != c.Direction {
				t.Errorf("Direction: %v", dir)
			}
		})
	}
}

func Test_arrowWithTailOnly(t *testing.T) {
	a := NewArrow(50, 50, 100, 50)
	a.Tail = NewCircle(3)
	a.Head = nil
	saveAs(t, a, "testdata/arrow_with_tail.svg")
}

func Test_arrowHasBothTailAndHead(t *testing.T) {
	a := NewArrow(50, 50, 100, 50)
	a.Tail = NewDiamond()
	a.Tail.SetX(20)
	a.Tail.SetY(24)
	a.Head = NewDiamond()
	a.Head.SetX(16)
	a.Head.SetY(8)
	saveAs(t, a, "testdata/arrow_with_tail_and_head.svg")
}

func TestLine(t *testing.T) {
	testShape(t, NewLine(0, 0, 50, 50))
}

func TestArrow_AbsAngle(t *testing.T) {
	a := NewLine(0, 0, 10, 10)
	got := a.AbsAngle()
	assert := asserter.New(t)
	assert().Equals(got, 45)
}

func TestArrow_Angle(t *testing.T) {
	a := NewLine(0, 0, 10, 10)
	got := a.Angle()
	assert := asserter.New(t)
	assert().Equals(got, 45)
}

// ----------------------------------------

func Test_ArrowBetweenShapes(t *testing.T) {
	a := NewRecord("A")
	a.SetX(10)
	a.SetY(100)
	b := NewRecord("B")
	b.SetX(80)
	b.SetY(40)

	arrow := NewArrowBetween(a, b)
	svg := &draw.SVG{}
	svg.SetSize(200, 200)
	svg.Append(arrow, a, b)
	label := NewLabel(fmt.Sprintf("Angle: %v", arrow.absAngle()))
	label.SetX(100)
	label.SetY(80)

	svg.Append(label)
	writeSvgTo(t, "testdata/arrow_between_shapes.svg", svg)
}

func saveAs(t *testing.T, line *Line, filename string) {
	t.Helper()
	d := &draw.SVG{}
	d.SetSize(100, 100)
	d.Append(line)

	writeSvgTo(t, filename, d)
}

func writeSvgTo(t *testing.T, filename string, svg *draw.SVG) {
	t.Helper()
	fh, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	style := draw.NewStyle()
	style.SetOutput(fh)
	svg.WriteSVG(&style)
	fh.Close()
}
