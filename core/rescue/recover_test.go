package rescue

import (
	"go-zero-my/core/logx"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	logx.Disable()
}

func TestRescue(t *testing.T) {
	var count int32
	assert.NotPanics(t, func() {
		defer Recover(func() {
			atomic.AddInt32(&count, 2)
		}, func() {
			atomic.AddInt32(&count, 3)
		})

		panic("hello")
	})
	assert.Equal(t, int32(5), atomic.LoadInt32(&count))
}