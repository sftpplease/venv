package venv

import (
	"flag"
	"fmt"
	"io"
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

type Os struct {
	Args   []string
	Stdin  io.ReadCloser
	Stdout io.WriteCloser
	Stderr io.WriteCloser
	Open   func(path string) (File, error)
	Stat   func(path string) (os.FileInfo, error)
	Exit   func(code int)
}

type Env struct {
	Os   *Os
	Flag *Flag
}

type Flag struct {
	vos     *Os
	flagSet *flag.FlagSet
}

func (f *Flag) Parse() {
	f.flagSet.Parse(f.vos.Args[1:])
}

func (f *Flag) PrintDefaults() {
	f.flagSet.PrintDefaults()
}

func (f *Flag) Args() []string {
	return f.flagSet.Args()
}

func (f *Flag) Bool(name string, value bool, usage string) *bool {
	return f.flagSet.Bool(name, value, usage)
}

func (f *Flag) Uint(name string, value uint, usage string) *uint {
	return f.flagSet.Uint(name, value, usage)
}

func NewFlag(vos *Os) *Flag {
	flagSet := flag.NewFlagSet(vos.Args[0], flag.ContinueOnError)
	flagSet.SetOutput(vos.Stderr)
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flagSet.PrintDefaults()
		vos.Exit(2)
	}

	return &Flag{
		vos:     vos,
		flagSet: flagSet,
	}
}
