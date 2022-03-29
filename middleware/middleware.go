package middleware

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// init a struct to handle the proemtheus counters
type PrometheusMiddleware struct {
	opsProcessed *prometheus.CounterVec
}

// initiate prometheus middleware
func NewPrometheusMiddleware() *PrometheusMiddleware {
	return &PrometheusMiddleware{
		opsProcessed: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "goprom_processed_ops_total",
			Help: "The total number of processed events",
		}, []string{"method", "path", "statuscode"}),
	}
}

// collect metrics from request
func (pm *PrometheusMiddleware) Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		pm.opsProcessed.With(
			prometheus.Labels{
				"method":     r.Method,
				"path":       r.RequestURI,
				"statuscode": "200",
			},
		).Inc()
	})
}
