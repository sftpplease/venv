package passthrough

import (
	"os"

	"github.com/sftpplease/venv"
)

func open(path string) (venv.File, error) {
	f, err := os.Open(path)
	return f, err
}

func PassthroughOS() *venv.Os {
	return &venv.Os{
		Args:   os.Args,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Open:   open,
		Stat:   os.Stat,
		Exit:   os.Exit,
	}
}
