package design

import (
	"fmt"
	"io"
	"time"

	"github.com/gregoryv/draw/shape"
)

// NewGanttChartFrom returns a GanttChart spanning days from the given
// date.  Panics if date cannot be resolved.
func NewGanttChartFrom(days int, from DateStr) *GanttChart {
	t := from.Time()
	return NewGanttChart(days, t)
}

// NewGanttChart returns a chart showing days from optional
// start time. If no start is given, time.Now() is used.
func NewGanttChart(days int, start ...time.Time) *GanttChart {
	d := &GanttChart{
		Diagram: NewDiagram(),
		start:   time.Now(),
		days:    days,
		tasks:   make([]*Task, 0),
		padLeft: 16,
		padTop:  10,
		Mark:    time.Now(),
	}
	if len(start) > 0 {
		d.start = start[0]
	}
	return d
}

type GanttChart struct {
	Diagram
	start time.Time
	days  int
	tasks []*Task

	padLeft, padTop int

	// Set a marker at this date.
	Mark time.Time
}

func (d *GanttChart) MarkDate(yyyymmdd DateStr) {
	d.Mark = yyyymmdd.Time()
}

// isToday returns true if time.Now matches start + ndays
func (d *GanttChart) isToday(ndays int) bool {
	t := d.start.AddDate(0, 0, ndays)
	return t.Year() == d.Mark.Year() &&
		t.YearDay() == d.Mark.YearDay()
}

// Add new task. Default color is green.
func (d *GanttChart) Add(txt string, offset, days int) *Task {
	task := NewTask(txt, offset, days)
	d.tasks = append(d.tasks, task)
	return task
}

// NewTask returns a green task.
func NewTask(txt string, offset, days int) *Task {
	return &Task{
		txt:    txt,
		offset: offset,
		days:   days,
		class:  "span-green",
	}
}

// Task is the colorized span of a gantt chart.
type Task struct {
	txt          string
	offset, days int
	class        string
}

// Red sets class of task to span-red
func (t *Task) Red() { t.class = "span-red" }

// Blue sets class of task to span-blue
func (t *Task) Blue() { t.class = "span-blue" }

func (d *GanttChart) WriteSvg(w io.Writer) error {
	now := d.start
	year := shape.NewLabel(fmt.Sprintf("%v", now.Year()))
	d.Place(year).At(d.padLeft, d.padTop)
	offset := d.padLeft + d.taskWidth()

	var lastDay *shape.Label
	columns := make([]*shape.Label, d.days)
	for i := 0; i < d.days; i++ {
		day := now.Day()
		col := newCol(day)
		columns[i] = col
		if now.Weekday() == time.Saturday {
			bg := shape.NewRect("")
			bg.SetClass("weekend")
			bg.SetWidth(col.Width()*2 + 8)
			bg.SetHeight(len(d.tasks)*col.Font.LineHeight + d.padTop + d.Diagram.Font.LineHeight)
			d.Place(bg).RightOf(lastDay, 4)
			shape.Move(bg, -2, 4)
		}
		if i == 0 {
			d.Place(col).Below(year, 4)
			col.SetX(offset)
		} else {
			d.Place(col).RightOf(lastDay, 4)
		}
		if day == 1 {
			label := shape.NewLabel(now.Month().String())
			d.Place(label).Above(col, 4)
		}
		if d.isToday(i) {
			x, y := col.Position()
			mark := shape.NewLine(x, y, x+10, y)
			d.Place(mark)
		}
		lastDay = col
		now = now.AddDate(0, 0, 1)
	}

	var lastTask *shape.Label
	for i, t := range d.tasks {
		label := shape.NewLabel(t.txt)
		if i == 0 {
			d.Place(label).Below(lastDay, 4)
			d.VAlignLeft(year, label)
		} else {
			d.Place(label).Below(lastTask, 4)
		}
		lastTask = label

		rect := shape.NewRect("")
		col := columns[t.offset]
		var w int
		for i := t.offset; i < t.offset+t.days; i++ {
			w += columns[i].Width() + 4
		}
		rect.SetWidth(w - 4)
		rect.SetHeight(d.Diagram.Font.Height)
		rect.SetClass(t.class)

		d.Place(rect).Below(col, 4)
		d.HAlignCenter(label, rect)
	}
	return d.Diagram.WriteSvg(w)
}

func newCol(day int) *shape.Label {
	col := shape.NewLabel(fmt.Sprintf("%02v", day))
	col.Font.Height = 10
	return col
}

func (d *GanttChart) SaveAs(filename string) error {
	return saveAs(d, d.Diagram.Style, filename)
}

func (d *GanttChart) taskWidth() int {
	x := 0
	for _, t := range d.tasks {
		w := d.Diagram.Font.TextWidth(t.txt)
		if w > x {
			x = w
		}
	}
	return x + d.padLeft
}

// DateStr has the format of yyyymmdd
type DateStr string

func (s DateStr) Time() time.Time {
	var (
		year  string
		month string
		day   string
	)
	switch len(s) {
	case 8:
		year = string(s[:4])
		month = string(s[4:6])
		day = string(s[6:])
	default:
		panic(fmt.Sprintf("unexpeced format yyyymmdd: %s", s))
	}
	str := fmt.Sprintf("%s-%02s-%02sT00:00:00.000Z", year, month, day)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		panic(err)
	}
	return t
}
