package sysx

import "go.uber.org/automaxprocs/maxprocs"

// Automatically set GOMAXPROCS to match Linux container CPU quota.
// 给容器设置 maxprocs, 否则获取到的是宿主机的maxcpu
func init() {
	maxprocs.Set(maxprocs.Logger(nil))
}
