package design_test

import (
	"database/sql"
	"testing"

	"github.com/drewjya/draw/design"
	"github.com/drewjya/draw/internal/app"
)

func TestExample(t *testing.T) {
	ExampleSequenceDiagram()
}

func ExampleSequenceDiagram() {
	var (
		d   = design.NewSequenceDiagram()
		cli = d.AddStruct(app.Client{})
		srv = d.AddStruct(app.Server{})
		db  = d.AddStruct(sql.DB{})
	)
	d.Group(cli, srv, "Public https", "blue") // default colors classes red, green, blue
	d.Group(srv, db, "Private rpc via Gob", "red")

	d.Link(cli, srv, "connect()")
	d.Link(srv, db, "SELECT").Class = "highlight"
	d.Link(db, srv, "Rows")
	d.Link(srv, srv, "Transform to view model").Class = "highlight"
	d.Skip()
	d.Return(srv, cli, "Send HTML")
	d.SaveAs("img/app_sequence_diagram.svg")
}
