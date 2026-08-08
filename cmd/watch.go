package cmd

import (
	"time"

	"github.com/devguilhrm/sysmonitor/internal/collector"
	"github.com/devguilhrm/sysmonitor/pkg/display"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Monitora CPU e Memória em tempo real com dashboard visual",
	Run: func(cmd *cobra.Command, args []string) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			memStats, _ := collector.GetMemStats()
			cpuStats, _ := collector.GetCPUStats()

			display.RenderDashboard(memStats, cpuStats)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
