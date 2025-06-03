package service

import (
	"github.com/meta-atman/metarpc/core/load"
	"github.com/meta-atman/metarpc/core/logger"
	"github.com/meta-atman/metarpc/core/proc"
	"github.com/meta-atman/metarpc/core/prometheus"
	"github.com/meta-atman/metarpc/core/stat"
	"github.com/meta-atman/metarpc/core/trace"
	"github.com/meta-atman/metarpc/internal/devserver"
)

const (
	// DevMode means development mode.
	DevMode = "dev"
	// TestMode means test mode.
	TestMode = "test"
	// RtMode means regression test mode.
	RtMode = "rt"
	// PreMode means pre-release mode.
	PreMode = "pre"
	// ProMode means production mode.
	ProMode = "pro"
)

type (
	// DevServerConfig is type alias for devserver.Config
	DevServerConfig = devserver.Config

	// A ServiceConf is a service config.
	ServiceConf struct {
		Name       string
		Log        logger.LogConf
		Mode       string `json:",default=pro,options=dev|test|rt|pre|pro"`
		MetricsUrl string `json:",optional"`
		// Deprecated: please use DevServer
		Prometheus prometheus.Config `json:",optional"`
		Telemetry  trace.Config      `json:",optional"`
		DevServer  DevServerConfig   `json:",optional"`
		Shutdown   proc.ShutdownConf `json:",optional"`
	}
)

// MustSetUp sets up the service, exits on error.
func (sc ServiceConf) MustSetUp() {
	logger.Must(sc.SetUp())
}

// SetUp sets up the service.
func (sc ServiceConf) SetUp() error {
	if len(sc.Log.ServiceName) == 0 {
		sc.Log.ServiceName = sc.Name
	}
	if err := logger.SetUp(sc.Log); err != nil {
		return err
	}

	sc.initMode()
	prometheus.StartAgent(sc.Prometheus)

	if len(sc.Telemetry.Name) == 0 {
		sc.Telemetry.Name = sc.Name
	}
	trace.StartAgent(sc.Telemetry)
	proc.Setup(sc.Shutdown)
	proc.AddShutdownListener(func() {
		trace.StopAgent()
	})

	if len(sc.MetricsUrl) > 0 {
		stat.SetReportWriter(stat.NewRemoteWriter(sc.MetricsUrl))
	}
	devserver.StartAgent(sc.DevServer)

	return nil
}

func (sc ServiceConf) initMode() {
	switch sc.Mode {
	case DevMode, TestMode, RtMode, PreMode:
		load.Disable()
		stat.SetReporter(nil)
	}
}
