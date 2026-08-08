package display

import (
	"fmt"

	"github.com/devguilhrm/sysmonitor/internal/models"
)

func RenderDashboard(mem models.MemStats, cpu models.CPUStats) {
	fmt.Print("\033[H\033[?25l")

	cyan := "\033[36m"
	green := "\033[32m"
	yellow := "\033[33m"
	red := "\033[31m"
	reset := "\033[0m"

	// Desenha o painel em bloco fixo
	fmt.Printf("%s┌───────────────────────────────────────────────┐%s\n", cyan, reset)
	fmt.Printf("%s│             SYS MONITOR - DASHBOARD           │%s\n", cyan, reset)
	fmt.Printf("%s└───────────────────────────────────────────────┘%s\n", cyan, reset)

	// Seção Memória RAM
	fmt.Printf("\n%s[ MEMÓRIA RAM ]%s\n", green, reset)
	fmt.Printf(" ├─ Total:       %.2f GB                     \n", mem.TotalGB)
	fmt.Printf(" ├─ Em Uso:      %.2f GB                     \n", mem.UsedGB)
	fmt.Printf(" └─ Porcentagem: %.1f%%                      \n", mem.PercentUse)

	// Seção Processador (CPU)
	fmt.Printf("\n%s[ PROCESSADOR ]%s\n", green, reset)
	fmt.Printf(" ├─ Modelo:      %s                          \n", truncateString(cpu.ModelName, 35))
	fmt.Printf(" ├─ Núcleos:     %d Físicos / %d Lógicos     \n", cpu.PhysicalCores, cpu.LogicalCores)

	corUso := green
	if cpu.TotalPercent > 70 {
		corUso = yellow
	}
	if cpu.TotalPercent > 90 {
		corUso = red
	}
	fmt.Printf(" └─ Uso Total:   %s%.2f%%%s                     \n", corUso, cpu.TotalPercent, reset)

	fmt.Println("\n-------------------------------------------------")
	fmt.Println(" Pressione [Ctrl + C] para sair.                 ")
}

func truncateString(str string, maxLen int) string {
	if len(str) > maxLen {
		return str[:maxLen] + "..."
	}
	return str
}
