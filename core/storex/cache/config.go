package cache

import "go-zero-my/core/storex/redis"
type (
	// A ClusterConf is the config of a redis cluster that used as cache.
	ClusterConf []NodeConf

	// A NodeConf is the config of a redis node that used as cache.
	NodeConf struct {
		redis.RedisConf
		Weight int `json:",default=100"`
	}
)