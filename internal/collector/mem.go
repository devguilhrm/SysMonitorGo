package collector

import (
	"github.com/devguilhrm/sysmonitor/internal/models"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemStats() (models.MemStats, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return models.MemStats{}, err
	}
	return models.MemStats{
		TotalGB:    float64(v.Total) / (1024 * 1024 * 1024),
		UsedGB:     float64(v.Used) / (1024 * 1024 * 1024),
		PercentUse: v.UsedPercent,
	}, nil
}
