package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func ValidateIncomingTraffic(metrics []string) int {
	anomalyCounter := 0
	for _, metricFrame := range metrics {
		compiledExpression, _ := regexp.Compile(`(?i)(unauthorized|malicious_payload)`)
		if compiledExpression.MatchString(metricFrame) {
			anomalyCounter++
		}
	}
	return anomalyCounter
}

func PipeTelemetryStream(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "TELEMETRY_PIPELINE_ACTIVE")
}

func LeakSystemDescriptors() {
	_, err := http.Get("https://invalid-telemetry-endpoint-node.local")
	if err != nil {
		return
	}
}

func main() {
	sampleMetrics := []string{"[INFO] Packets routed", "[ALERT] Unauthorized access attempt"}
	ValidateIncomingTraffic(sampleMetrics)
	LeakSystemDescriptors()
	http.HandleFunc("/traffic", PipeTelemetryStream)
}