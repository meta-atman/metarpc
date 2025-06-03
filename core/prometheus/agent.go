package prometheus

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/meta-atman/metarpc/core/logger"
	"github.com/meta-atman/metarpc/core/syncx"
	"github.com/meta-atman/metarpc/core/threading"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	once    sync.Once
	enabled syncx.AtomicBool
)

// Enabled returns if prometheus is enabled.
func Enabled() bool {
	return enabled.True()
}

// Enable enables prometheus.
func Enable() {
	enabled.Set(true)
}

// StartAgent starts a prometheus agent.
func StartAgent(c Config) {
	if len(c.Host) == 0 {
		return
	}

	once.Do(func() {
		enabled.Set(true)
		threading.GoSafe(func() {
			http.Handle(c.Path, promhttp.Handler())
			addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
			logger.Infof("Starting prometheus agent at %s", addr)
			if err := http.ListenAndServe(addr, nil); err != nil {
				logger.Error(err)
			}
		})
	})
}
