package display

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/devguilhrm/sysmonitor/internal/models"
)

func RenderMemTable(stats models.MemStats) {
	// Inicializa o tabwriter para alinhar as colunas perfeitamente no terminal
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

	// Imprime o cabeçalho principal
	fmt.Fprintln(w, "=========================================")
	fmt.Fprintln(w, "       ESTATÍSTICAS DE MEMÓRIA RAM       ")
	fmt.Fprintln(w, "=========================================")
	fmt.Fprintln(w, "MÉTRICA\tVALOR")
	fmt.Fprintln(w, "-------\t-----")

	// Imprime as linhas de dados dinâmicos
	fmt.Fprintf(w, "Total (GB)\t%.2f GB\n", stats.TotalGB)
	fmt.Fprintf(w, "Em Uso (GB)\t%.2f GB\n", stats.UsedGB)

	// Adiciona um alerta visual simples se o uso estiver muito alto
	if stats.PercentUse > 85.0 {
		fmt.Fprintf(w, "Uso (%%)\t%.2f%% [ALERTA: USO ALTO]\n", stats.PercentUse)
	} else {
		fmt.Fprintf(w, "Uso (%%)\t%.2f%%\n", stats.PercentUse)
	}

	fmt.Fprintln(w, "=========================================")

	// Flush aplica a formatação e joga tudo na tela
	w.Flush()
}
