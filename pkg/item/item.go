package item

import "github.com/spf13/pflag"

type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func FlagSet() *pflag.FlagSet {
	fs := pflag.NewFlagSet("item", pflag.ExitOnError)
	fs.IntVar(&itemStore.MaxKeyLen, "max-key-len", 10, "maximum length of redis keys")
	fs.IntVar(&itemStore.MinKeyLen, "min-key-len", 5, "minimum length of redis keys")
	fs.IntVar(&itemStore.MaxValueLen, "max-value-len", 100, "maximum length of redis values")
	fs.IntVar(&itemStore.MinValueLen, "min-value-len", 30, "minimum length of redis values")
	return fs
}
