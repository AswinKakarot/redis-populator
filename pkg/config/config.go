package config

import "github.com/spf13/pflag"

type AppConfig struct {
	ReadDuration,
	WriteDuration,
	ReadInterval,
	WriteInterval int
}

var (
	cfg *AppConfig
)

func FlagSet() *pflag.FlagSet {
	fs := pflag.NewFlagSet("app-config", pflag.ExitOnError)
	fs.IntVar(&cfg.ReadDuration, "read-duration", 10, "duration to run read operation on the cluster in seconds.")
	fs.IntVar(&cfg.WriteDuration, "write-duration", 10, "duration to run write operation on the cluster in seconds.")
	fs.IntVar(&cfg.ReadInterval, "read-interval", 2, "interval between consecutive read operations on the cluster in seconds.")
	fs.IntVar(&cfg.WriteInterval, "write-interval", 1, "interval between consecutive write operations on the cluster in seconds.")
	return fs
}

func GetAppConfig() *AppConfig {
	return cfg
}

func init() {
	cfg = &AppConfig{}
}
