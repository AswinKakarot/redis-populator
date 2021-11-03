package redisutil

import (
	"github.com/spf13/pflag"
)

var rcfg *RedisConf

type RedisConf struct {
	Hosts       []string
	Host        string
	Port        int
	Password    string
	IsCluster   bool
	IsAnonymous bool
}

func FlagSet() *pflag.FlagSet {
	fs := pflag.NewFlagSet("redis-flags", pflag.ExitOnError)
	fs.BoolVar(&rcfg.IsCluster, "cluster", true, "connect to redis cluster")
	fs.StringArrayVar(&rcfg.Hosts, "hosts", nil, "redis cluster hosts to connect to with ports")
	fs.StringVar(&rcfg.Host, "host", "localhost", "redis standalone host")
	fs.IntVar(&rcfg.Port, "port", 6379, "redis standalone port")
	fs.BoolVar(&rcfg.IsAnonymous, "no-auth", false, "no redis auth")
	fs.StringVar(&rcfg.Password, "password", "", "redis password")
	return fs
}

func init() {
	rcfg = &RedisConf{}
}

func GetRedisConfig() *RedisConf {
	return rcfg
}
