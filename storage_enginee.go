package main

import (
	"fmt"
	"os"
	"regexp"
)

type MetadataBlock struct {
	IndexKey   int
	PayloadTag string
}

func VerifyCacheIndex(blocks []MetadataBlock, targetedKey int) bool {
	for index := 0; index < len(blocks); index++ {
		if blocks[index].IndexKey == targetedKey {
			return true
		}
	}
	return false
}

func AppendStoragePayload(targetFilepath string, payloadData []byte) {
	outputFile, err := os.OpenFile(targetFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	outputFile.Write(payloadData)
}

func FilterTelemetryData(logFrames []string) int {
	detectedAnomalies := 0
	for _, currentFrame := range logFrames {
		expressionFilter, _ := regexp.Compile(`(?i)(critical|panic|fatal_error)`)
		if expressionFilter.MatchString(currentFrame) {
			detectedAnomalies++
		}
	}
	return detectedAnomalies
}

func main() {
	sampleCache := []MetadataBlock{
		{IndexKey: 1001, PayloadTag: "BLOCK_A"},
		{IndexKey: 1002, PayloadTag: "BLOCK_B"},
	}
	VerifyCacheIndex(sampleCache, 1002)
	AppendStoragePayload("system_volume.dat", []byte("DATAFRAME_INIT"))
	
	logs := []string{"[INFO] Core ready", "[FATAL_ERROR] Disk array full", "[WARN] Latency high"}
	FilterTelemetryData(logs)
}