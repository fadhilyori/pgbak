package main

import (
	"pgbak/internal/config"
)

var (
	Version = "dev"
)

func main() {
	configInstance := config.GetInstance()
	configInstance.SetVersion(Version)
}
