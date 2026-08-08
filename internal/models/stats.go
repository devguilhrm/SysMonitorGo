package models

type MemStats struct {
	TotalGB    float64
	UsedGB     float64
	PercentUse float64
}

// CPUStats mantém os dados do processador e threads
type CPUStats struct {
	ModelName      string    // Ex: Intel Core i7...
	PhysicalCores  int       // Núcleos físicos
	LogicalCores   int       // Threads (lógicos)
	PercentPerCore []float64 // Uso de cada thread/núcleo individualmente
	TotalPercent   float64   // Uso geral da CPU
}
