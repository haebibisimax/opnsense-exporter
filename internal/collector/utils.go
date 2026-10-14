package collector

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

func buildPrometheusDesc(subsystem, name, help string, labels []string) *prometheus.Desc {
	if labels != nil {
		labels = append(labels, instanceLabelName)
	} else {
		labels = []string{instanceLabelName}
	}

	return prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, name),
		help,
		labels,
		nil,
	)
}

// parseStringToBool parses a string value to a bool value.
//
// An empty value is false. The OPNsense API omits fields rather than sending a
// falsy one, and the previous implementation ("true unless it is 0") turned a
// missing field into an enabled service in the exported metric.
func parseStringToBool(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "", "0", "false", "no", "off", "disabled":
		return false
	default:
		return true
	}
}
