package discov

import (
	"sync"
	"testing"

	"github.com/meta-atman/metarpc/core/discov/internal"
	"github.com/stretchr/testify/assert"
)

var mockLock sync.Mutex

func setMockClient(cli internal.EtcdClient) func() {
	mockLock.Lock()
	internal.NewClient = func([]string) (internal.EtcdClient, error) {
		return cli, nil
	}
	return func() {
		internal.NewClient = internal.DialClient
		mockLock.Unlock()
	}
}

func TestExtract(t *testing.T) {
	id, ok := extractId("key/123/val")
	assert.True(t, ok)
	assert.Equal(t, "123", id)

	_, ok = extract("any", -1)
	assert.False(t, ok)

	_, ok = extract("any", 10)
	assert.False(t, ok)
}

func TestMakeKey(t *testing.T) {
	assert.Equal(t, "key/123", makeEtcdKey("key", 123))
}
