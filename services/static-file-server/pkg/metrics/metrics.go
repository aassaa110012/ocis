package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// Namespace defines the namespace for the defines metrics.
	Namespace = "ocis"

	// Subsystem defines the subsystem for the defines metrics.
	Subsystem = "static-file-server"
)

// Metrics defines the available metrics of this service.
type Metrics struct {
	BuildInfo *prometheus.GaugeVec
}

// New initializes the available metrics.
func New() *Metrics {
	m := &Metrics{
		BuildInfo: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "build_info",
			Help:      "Build Information",
		}, []string{"version"}),
	}

	_ = prometheus.Register(
		m.BuildInfo,
	)

	return m
}
