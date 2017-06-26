package venv

import (
	"os"
)

type File interface {
	Read(buf []byte) (int, error)
	Readdir(n int) ([]os.FileInfo, error)
	Write(buf []byte) (int, error)
	WriteAt(buf []byte, off int64) (int, error)
	Stat() (os.FileInfo, error)
	Close() error
}

type VOS interface {
	Open(path string) (File, error)
	Stat(path string) (os.FileInfo, error)
	Exit(code int)
}
