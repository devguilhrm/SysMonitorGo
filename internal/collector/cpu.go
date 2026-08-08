package collector

import (
	"time"

	"github.com/devguilhrm/sysmonitor/internal/models"
	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCPUStats() (models.CPUStats, error) {
	info, err := cpu.Info()
	if err != nil {
		return models.CPUStats{}, err
	}

	physicalCores, err := cpu.Counts(false)
	if err != nil {
		physicalCores = 0
	}

	logicalCores, err := cpu.Counts(true)
	if err != nil {
		logicalCores = 0
	}

	modelName := "Desconhecido"
	if len(info) > 0 {
		modelName = info[0].ModelName
	}

	percentages, err := cpu.Percent(200*time.Millisecond, false)
	totalPercent := 0.0
	if err == nil && len(percentages) > 0 {
		totalPercent = percentages[0]
	}

	perCore, err := cpu.Percent(0, true)
	if err != nil {
		perCore = []float64{}
	}

	return models.CPUStats{
		ModelName:      modelName,
		PhysicalCores:  physicalCores,
		LogicalCores:   logicalCores,
		PercentPerCore: perCore,
		TotalPercent:   totalPercent,
	}, nil
}
