package redisutil

import (
	"context"
	"fmt"

	"github.com/mediocregopher/radix/v4"
)

type Standalone struct {
	client radix.Client
}

func NewStandalone() (o *Standalone, err error) {
	o = new(Standalone)
	var dialer radix.Dialer
	if !rcfg.IsAnonymous {
		dialer = radix.Dialer{
			AuthPass: rcfg.Password,
		}
	}
	cfg := radix.PoolConfig{
		Dialer: dialer,
	}
	o.client, err = cfg.New(context.TODO(), "tcp", fmt.Sprintf("%s:%d", rcfg.Host, rcfg.Port))
	return
}

func (c *Standalone) Write(key, value string) error {
	return c.client.Do(context.TODO(), radix.Cmd(nil, "SET", key, value))
}

func (c *Standalone) Read(key string) (value string, err error) {
	err = c.client.Do(context.TODO(), radix.Cmd(&value, "GET", key))
	return
}
