package threading

import (
	"bytes"
	"context"
	"io"
	"log"
	"testing"

	"github.com/meta-atman/metarpc/core/cast"
	"github.com/meta-atman/metarpc/core/logger"
	"github.com/stretchr/testify/assert"
)

func TestRoutineId(t *testing.T) {
	assert.True(t, RoutineId() > 0)
}

func TestRunSafe(t *testing.T) {
	log.SetOutput(io.Discard)

	i := 0

	defer func() {
		assert.Equal(t, 1, i)
	}()

	ch := make(chan cast.PlaceholderType)
	go RunSafe(func() {
		defer func() {
			ch <- cast.Placeholder
		}()

		panic("panic")
	})

	<-ch
	i++
}

func TestRunSafeCtx(t *testing.T) {
	var buf bytes.Buffer
	logger.SetWriter(logger.NewWriter(&buf))
	ctx := context.Background()
	ch := make(chan cast.PlaceholderType)

	i := 0

	defer func() {
		assert.Equal(t, 1, i)
	}()

	go RunSafeCtx(ctx, func() {
		defer func() {
			ch <- cast.Placeholder
		}()

		panic("panic")
	})

	<-ch
	i++
}

func TestGoSafeCtx(t *testing.T) {
	var buf bytes.Buffer
	logger.SetWriter(logger.NewWriter(&buf))
	ctx := context.Background()
	ch := make(chan cast.PlaceholderType)

	i := 0

	defer func() {
		assert.Equal(t, 1, i)
	}()

	GoSafeCtx(ctx, func() {
		defer func() {
			ch <- cast.Placeholder
		}()

		panic("panic")
	})

	<-ch
	i++
}
