package redis

import (
	"testing"

	"github.com/meta-atman/metarpc/core/logger"
	"github.com/stretchr/testify/assert"
)

func TestScriptCache(t *testing.T) {
	logger.Disable()

	cache := GetScriptCache()
	cache.SetSha("foo", "bar")
	cache.SetSha("bla", "blabla")
	bar, ok := cache.GetSha("foo")
	assert.True(t, ok)
	assert.Equal(t, "bar", bar)
}
