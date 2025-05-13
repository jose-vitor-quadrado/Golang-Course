package main

import (
	"io"
	"os"
)

// Keep interface small, defining only minimum behavior
type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

func main() {

}
