package main

import (
	_ "embed"
	"github.com/howardjohn/shgo/pkg"
)

//go:embed script.sh
var script []byte

func main() {
	if err := pkg.Exec("hello-world", script); err != nil {
		panic(err)
	}
}
