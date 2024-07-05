package shgo

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"syscall"
)

func Exec(name string, executable []byte) error {
	fp, err := MemFile(name, executable)
	if err != nil {
		return err
	}

	args := []string{name}
	env := os.Environ()
	return syscall.Exec(fp, args, env)
}

// MemFile takes a file name used (mostly for debugging), and the contents the file should contain, and returns the file name.
func MemFile(name string, b []byte) (string, error) {
	fd, err := MemFd(name, b)
	if err != nil {
		return "", err
	}
	// filepath to our newly created in-memory file descriptor
	fp := fmt.Sprintf("/proc/self/fd/%d", fd)
	return fp, nil
}

// MemFd takes a file name used (mostly for debugging), and the contents the file should contain, and returns the file descriptor.
func MemFd(name string, b []byte) (int, error) {
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
