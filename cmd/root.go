package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sysmonitor",
	Short: "Uma CLI rápida para monitoramento de sistema",
	Long:  `Sysmonitor é uma ferramenta de linha de comando em Go para verificar o uso de RAM, CPU, Disco e Rede.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Erro ao executar o comando:", err)
		os.Exit(1)
	}
}
