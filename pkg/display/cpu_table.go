package display

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/devguilhrm/sysmonitor/internal/models"
)

// RenderCPUTable desenha as estatísticas do processador no terminal
func RenderCPUTable(stats models.CPUStats) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

	fmt.Fprintln(w, "=========================================")
	fmt.Fprintln(w, "          ESTATÍSTICAS DE CPU            ")
	fmt.Fprintln(w, "=========================================")
	fmt.Fprintln(w, "MÉTRICA\tVALOR")
	fmt.Fprintln(w, "-------\t-----")

	fmt.Fprintf(w, "Modelo\t%s\n", stats.ModelName)
	fmt.Fprintf(w, "Núcleos Físicos\t%d\n", stats.PhysicalCores)
	fmt.Fprintf(w, "Threads (Lógicos)\t%d\n", stats.LogicalCores)
	fmt.Fprintf(w, "Uso Total (%%)\t%.2f%%\n", stats.TotalPercent)

	fmt.Fprintln(w, "-----------------------------------------")
	fmt.Fprintln(w, "USO POR THREAD (NÚCLEO)")
	fmt.Fprintln(w, "-----------------------------------------")

	for i, usage := range stats.PercentPerCore {
		fmt.Fprintf(w, "Thread %d\t%.2f%%\n", i+1, usage)
	}

	fmt.Fprintln(w, "=========================================")
	w.Flush()
}
