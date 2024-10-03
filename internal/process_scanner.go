package internal

import (
	"log"
	"sort"
	"time"

	"github.com/shirou/gopsutil/process"
)

const MB = 1000000

type Process struct {
	Id     int
	Name   string
	Memory float32
	CPU    float32
}

// Scanner: Lista os processos em execucao com informacoes de Id, nome, %memoria, %cpu
func scanner() []*Process {
	var processList []*Process

	processes, err := process.Processes()
	if err != nil {
		log.Fatalf("Error in Scanner: %s", err)
	}

	for _, idx := range processes {
		id := idx.Pid

		name, err := idx.Name()
		if err != nil {
			name = "Desconhecido"
		}

		memInfo, err := idx.MemoryInfo()
		if err != nil {
			memInfo = &process.MemoryInfoStat{}
		}

		cpuPercent, err := idx.CPUPercent()
		if err != nil {
			cpuPercent = 0
		}

		processList = append(processList, &Process{
			Id:     int(id),
			Name:   name,
			Memory: float32(memInfo.RSS * (1024 * 1024)),
			CPU:    float32(cpuPercent),
		})
	}
	time.Sleep(100 * time.Millisecond)

	sort.Slice(processList, func(i, j int) bool {
		return processList[i].Memory > processList[j].Memory
	})

	return processList
}

func Run() []*Process {
	return scanner()
}
