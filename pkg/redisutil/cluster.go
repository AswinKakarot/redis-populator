package redisutil

import (
	"context"

	radix "github.com/mediocregopher/radix/v4"
)

type Cluster struct {
	client *radix.Cluster
}

func NewCluster() (*Cluster, error) {
	o := &Cluster{}
	var err error
	dialer := new(radix.Dialer)
	if !rcfg.IsAnonymous {
		dialer.AuthPass = rcfg.Password
	}
	cfg := &radix.ClusterConfig{
		PoolConfig: radix.PoolConfig{
			Dialer: *dialer,
		},
	}
	o.client, err = cfg.New(context.TODO(), rcfg.Hosts)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (c *Cluster) Write(key, value string) error {
	return c.client.Do(context.TODO(), radix.Cmd(nil, "SET", key, value))
}

func (c *Cluster) Read(key string) (value string, err error) {
	err = c.client.Do(context.TODO(), radix.Cmd(&value, "GET", key))
	return
}
