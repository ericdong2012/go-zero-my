package prof

import (
	"fmt"
	"testing"

	"go-zero-my/core/utils"
)

func TestProfiler(t *testing.T) {
	EnableProfiling()
	start := Start()
	fmt.Println(start)
	Report("foo", ProfilePoint{
		ElapsedTimer: utils.NewElapsedTimer(),
	})
}

func TestNullProfiler(t *testing.T) {
	p := newNullProfiler()
	p.Start()
	p.Report("foo", ProfilePoint{
		ElapsedTimer: utils.NewElapsedTimer(),
	})
}
