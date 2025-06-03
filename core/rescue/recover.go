package rescue

import (
	"context"
	"runtime/debug"

	"github.com/meta-atman/metarpc/core/logger"
)

// Recover is used with defer to do cleanup on panics.
// Use it like:
//
//	defer Recover(func() {})
func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		logger.ErrorStack(p)
	}
}

// RecoverCtx is used with defer to do cleanup on panics.
func RecoverCtx(ctx context.Context, cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		logger.WithContext(ctx).Errorf("%+v\n%s", p, debug.Stack())
	}
}
