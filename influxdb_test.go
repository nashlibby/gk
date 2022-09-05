package gk

import (
	"testing"
)

func getClient() *InfluxClient {
	return NewInfluxClient(InfluxSetting{
		Host:   "http://127.0.0.1:8086",
		Token:  "Rdd5Fq81Lm8sRiI-LAnIq-74NSYKFRs36T74v5MbLv99cduZRJ-sbfLLpMUZAxsGhUpNs74h86SEA2IU6ANzcw==",
		Org:    "gindow",
		Bucket: "supplier",
	})
}

func TestInfluxClient_Write(t *testing.T) {
	client := getClient()
	client.Write("api", map[string]string{
		"unit": "temperature",
	}, map[string]interface{}{
		"avg": 24.5,
		"max": 40,
		"min": 15,
	})
}
