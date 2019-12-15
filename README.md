[![Build Status](https://travis-ci.org/gregoryv/draw.svg?branch=master)](https://travis-ci.org/gregoryv/draw)
[![codecov](https://codecov.io/gh/gregoryv/draw/branch/master/graph/badge.svg)](https://codecov.io/gh/gregoryv/draw)
[![Maintainability](https://api.codeclimate.com/v1/badges/b0001c5ba7cd098b183d/maintainability)](https://codeclimate.com/github/gregoryv/draw/maintainability)

[draw](https://godoc.org/github.com/gregoryv/draw) - package for writing software design diagrams

- Cross platform
- No external dependencies
- SVG output

Program your diagrams and refactoring automatically updates them.
Take a look at the below examples and then browse the [showcase](./showcase/README.md) of golang standard packages.

## Sequence diagram

<img src="shape/design/img/app_sequence_diagram.svg">

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

## Activity diagram

<img src="shape/design/img/activity_diagram.svg">

Rendered by
[ExampleActivityDiagram](https://godoc.org/github.com/gregoryv/draw/shape/design/#example-ActivityDiagram)

## Class diagram

Class diagrams show relations between structs and
interfaces. Reflection includes fields and methods.

This diagram is rendered by
[ExampleClassDiagram](https://godoc.org/github.com/gregoryv/draw/shape/design/#example-ClassDiagram)

<img src="shape/design/img/class_example.svg" style="width: 500"/>

## Generic diagram

It should be easy to just add any extra shapes to any diagram when explaining a design.
This diagram is rendered by
[ExampleDiagram](https://godoc.org/github.com/gregoryv/draw/shape/design/#example-Diagram)

![](shape/design/img/diagram_example.svg)


## Grid layout

Simplifying placing shapes in a grid layout aligning different sizes of shapes.

![](shape/design/img/grid_layout.svg)


## TODO

- Labeled arrows
- Link to optional godoc service
- More shapes

WIP - major rewrites still going on
