package prof

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
	once.Do(func() {})
	assert.NotContains(t, generateReport(), "foo")
	report("foo", time.Second)
	s := generateReport()
	fmt.Println(s)
	assert.Contains(t, generateReport(), "foo")
	report("foo", time.Second)
}
