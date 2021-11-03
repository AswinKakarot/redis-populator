package util

import (
	"os"

	"github.com/spf13/pflag"
)

var (
	filepath string
	ls       = &LocalStore{}
)

type LocalStore struct {
	Fptr         *os.File
	isLocalStore bool
}

func FlagSet() *pflag.FlagSet {
	fs := pflag.NewFlagSet("local-store", pflag.ExitOnError)
	fs.BoolVar(&ls.isLocalStore, "localstore", false, "enable writing to file")
	fs.StringVar(&filepath, "data", "/data/outfile", "local storage location used for persistent storage.")
	return fs
}

func GetLocalStore() (*LocalStore, error) {
	fptr, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	ls.Fptr = fptr
	return ls, nil
}

func LocalStoreEnabled() bool {
	return ls.isLocalStore
}
