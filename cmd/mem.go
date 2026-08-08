package cmd

import (
	"github.com/devguilhrm/sysmonitor/internal/collector"
	"github.com/devguilhrm/sysmonitor/pkg/display"
	"github.com/spf13/cobra"
)

var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "Exibe estatísticas de uso da memória RAM",
	Run: func(cmd *cobra.Command, args []string) {
		stats, _ := collector.GetMemStats()
		display.RenderMemTable(stats)
	},
}

func init() {
	rootCmd.AddCommand(memCmd)
}
