package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type VirtualMemory struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

type Metrics struct {
	Hostname          string        `json:"hostname"`
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

func main() {
	serialized, err := json.Marshal(metrics)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(serialized))
}
