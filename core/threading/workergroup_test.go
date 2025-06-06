package threading

import (
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/meta-atman/metarpc/core/cast"
	"github.com/stretchr/testify/assert"
)

func TestWorkerGroup(t *testing.T) {
	m := make(map[string]cast.PlaceholderType)
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())
	group := NewWorkerGroup(func() {
		lock.Lock()
		m[fmt.Sprint(RoutineId())] = cast.Placeholder
		lock.Unlock()
		wg.Done()
	}, runtime.NumCPU())
	go group.Start()
	wg.Wait()
	assert.Equal(t, runtime.NumCPU(), len(m))
}
