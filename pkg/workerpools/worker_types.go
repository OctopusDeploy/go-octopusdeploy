package workerpools

type WorkerType string

const (
	WorkerTypeUbuntu1804      WorkerType = "Ubuntu1804"
	WorkerTypeUbuntuDefault   WorkerType = "UbuntuDefault"
	WorkerTypeWindows2016     WorkerType = "Windows2016"
	WorkerTypeWindows2019     WorkerType = "Windows2019"
	WorkerTypeUWindowsDefault WorkerType = "WindowsDefault"
)
