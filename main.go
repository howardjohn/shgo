package main

import (
	_ "embed"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

//go:embed script.sh
var script []byte

func main() {
	fd, err := memfile("shgo", script)
	if err != nil {
		panic(err)
	}

	// filepath to our newly created in-memory file descriptor
	fp := fmt.Sprintf("/proc/self/fd/%d", fd)

	// binary, err := exec.LookPath("bash")
	// if err != nil {
	// 	panic(err)
	// }
	// args := []string{"bash", "-c", fp}
	args := []string{fp}
	env := os.Environ()
	if err := syscall.Exec(fp, args, env); err != nil {
		panic(err)
	}
}

// memfile takes a file name used, and the byte slice
// containing data the file should contain.
//
// name does not need to be unique, as it's used only
// for debugging purposes.
//
// It is up to the caller to close the returned descriptor.
func memfile(name string, b []byte) (int, error) {
	fd, err := unix.MemfdCreate(name, 0)
	if err != nil {
		return 0, fmt.Errorf("MemfdCreate: %v", err)
	}

	err = unix.Ftruncate(fd, int64(len(b)))
	if err != nil {
		return 0, fmt.Errorf("Ftruncate: %v", err)
	}

	data, err := unix.Mmap(fd, 0, len(b), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return 0, fmt.Errorf("Mmap: %v", err)
	}

	copy(data, b)

	err = unix.Munmap(data)
	if err != nil {
		return 0, fmt.Errorf("Munmap: %v", err)
	}

	return fd, nil
}
