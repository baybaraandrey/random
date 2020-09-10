package main

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/shirou/gopsutil/mem"
)

type VirtualMemory struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

type Metrics struct {
	Hostname          string        `json:"hostname"`
	Time string `json:"time"`
	VirtualMemoryStat VirtualMemory `json:"virtual_memory_stat"`
}

var (
	metrics = Metrics{
		Hostname: "web",
		VirtualMemoryStat: VirtualMemory{
			Total:       3179569152,
			Free:        284233728,
			UsedPercent: 84.50,
		},
	}
)

func Gather() *Metrics {
	v, _ := mem.VirtualMemory()
	return &Metrics{
		Hostname: "web",
		VirtualMemoryStat: VirtualMemory{
			Total:       v.Total,
			Free:        v.Free,
			UsedPercent: v.UsedPercent,
		},
	}
}

func main() {

	for tick := range time.Tick(1 * time.Second) {
		metrics := Gather()
		metrics.Time = tick.String()
		serialized, err := json.Marshal(*metrics)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(serialized))
	}
}
