package main

import (
	_ "embed"
	"github.com/howardjohn/shgo"
)

//go:embed script.sh
var script []byte

func main() {
	if err := shgo.Exec("hello-world", script); err != nil {
		panic(err)
	}
}
