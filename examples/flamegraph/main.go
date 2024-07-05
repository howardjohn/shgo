package main

import (
	_ "embed"
	"github.com/howardjohn/shgo"
)

//go:embed flamegraph.pl
var script []byte

func main() {
	if err := shgo.Exec("flamegraph.pl", script); err != nil {
		panic(err)
	}
}
