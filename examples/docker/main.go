package main

import (
	_ "embed"
	"fmt"
	"github.com/howardjohn/shgo"
)

// Fetch with `docker pull hello-world; docker save hello-world | zstd > image.zst`
//go:embed image.zst
var image []byte

var loadScript = `#!/bin/bash
docker load -q -i %s > /dev/null
exec docker run --pull=never hello-world:latest
`

func main() {
	imageFile, err := shgo.MemFile("image", image)
	if err != nil {
		panic(err)
	}
	script := fmt.Sprintf(loadScript, imageFile)
	if err := shgo.Exec("docker", []byte(script)); err != nil {
		panic(err)
	}
}
