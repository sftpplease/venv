package passthrough

import (
	"os"

	"github.com/sftpplease/venv"
)

func open(path string) (venv.File, error) {
	f, err := os.Open(path)
	return f, err
}

func openFile(name string, flag int, perm os.FileMode) (venv.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	return f, err
}

func PassthroughOS() *venv.Os {
	return &venv.Os{
		Args:   os.Args,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Chmod: os.Chmod,
		Mkdir: os.Mkdir,
		Open:   open,
		OpenFile: openFile,
		Stat:   os.Stat,
		Exit:   os.Exit,
	}
}
