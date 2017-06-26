package passthrough

import (
	"os"

	"github.com/sftpplease/venv"
)

type Passthrough struct{}

func (pt *Passthrough) Open(path string) (venv.File, error) {
	f, err := os.Open(path)
	return f, err
}

func (pt *Passthrough) Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func (pt *Passthrough) Exit(code int) {
	os.Exit(code)
}

func New() venv.VOS {
	return &Passthrough{}
}
