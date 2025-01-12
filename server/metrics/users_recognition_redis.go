package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var userRecognitionRedisLabels = []string{"error_type"}

var (
	usersRecognitionRedisErrors *prometheus.CounterVec
)

func initUsersRecognitionRedis() {
	usersRecognitionRedisErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "eventnative",
		Subsystem: "users_recognition",
		Name:      "redis",
	}, userRecognitionRedisLabels)
}

func UserRecognitionRedisErrors(errorType string) {
	if Enabled {
		usersRecognitionRedisErrors.WithLabelValues(errorType).Inc()
	}
}
