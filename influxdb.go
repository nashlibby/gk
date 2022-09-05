package gk

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

type InfluxSetting struct {
	Host   string `json:"host"`
	Token  string `json:"token"`
	Org    string `json:"org"`
	Bucket string `json:"bucket"`
}

type InfluxClient struct {
	Client  influxdb2.Client `json:"client"`
	Setting InfluxSetting    `json:"setting"`
}

func NewInfluxClient(setting InfluxSetting) *InfluxClient {
	client := influxdb2.NewClient(setting.Host, setting.Token)
	return &InfluxClient{Client: client, Setting: setting}
}

func (i *InfluxClient) Write(measurement string, tags map[string]string, fields map[string]interface{}) {
	writeAPI := i.Client.WriteAPI(i.Setting.Org, i.Setting.Bucket)
	p := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	writeAPI.WritePoint(p)
	writeAPI.Flush()
}
