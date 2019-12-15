package main

import (
	"database/sql"

	"github.com/gregoryv/draw/internal/app"
	"github.com/gregoryv/draw/shape/design"
)

//go:generate go run .
func main() {
	var (
		d   = design.NewSequenceDiagram()
		cli = d.AddStruct(app.Client{})
		srv = d.AddStruct(app.Server{})
		db  = d.AddStruct(sql.DB{})
	)
	d.Link(cli, srv, "connect()")
	d.Link(srv, db, "SELECT").Class = "highlight"
	d.Link(db, srv, "Rows")
	d.Link(srv, srv, "Transform to view model").Class = "highlight"
	d.Link(srv, cli, "Send HTML")
	d.SaveAs("../../../shape/design/img/app_sequence_diagram.svg")
}
