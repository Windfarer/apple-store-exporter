package metric

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(
		followerMetric,
	)
}

var (
	followerMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "bb",
		Subsystem: "exporter",
		Name:      "follower",
		Help:      "",
	}, []string{"uid"})
)
