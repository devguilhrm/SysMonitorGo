package cmd

import (
	"github.com/devguilhrm/sysmonitor/internal/collector"
	"github.com/devguilhrm/sysmonitor/pkg/display"
	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Exibe estatísticas de uso do processador e threads",
	Run: func(cmd *cobra.Command, args []string) {
		stats, _ := collector.GetCPUStats()
		display.RenderCPUTable(stats)
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
}
