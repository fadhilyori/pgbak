package config

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"sync"
)

var (
	defaultConfigFilePath  = "/etc/default/pgbak"
	defaultHostname        = "postgres"
	defaultUsername        string
	defaultBackupDirPath   string
	defaultEnableGlobals   = false
	defaultEnablePlain     = true
	defaultEnableCustom    = false
	defaultSchemaOnlyList  = false
	defaultWeeksToKeep     = 12
	defaultDaysToKeep      = 0
	defaultDayOfWeekToKeep = 0

	instance *Config
	once     sync.Once
)

type Config struct {
	AppVersion      string
	CobraInstance   *cobra.Command
	Hostname        string
	Username        string
	Password        string
	BackupDirPath   string
	WeeksToKeep     int
	DaysToKeep      int
	DayOfWeekToKeep int
	EnableGlobals   bool
	EnablePlain     bool
	EnableCustom    bool
	SchemaOnlyList  bool
}

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
		instance.LoadConfig()
	})

	return instance
}

func (c *Config) SetVersion(version string) {
	c.AppVersion = version
	c.CobraInstance.Version = version
}

func (c *Config) LoadConfig() {
	c.CobraInstance = &cobra.Command{
		Use:   "pgbak [flags]",
		Short: "A tool for automated backup for PostgreSQL",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		PreRun: func(cmd *cobra.Command, args []string) {
			figure.NewFigure("pgbak", "larry3d", true).Print()
			fmt.Println(cmd.Version)
			fmt.Print("\n\n")
		},
	}

	flags := c.CobraInstance.Flags()

	filename := flags.StringP("config", "c", defaultConfigFilePath, "Specifies the config file path")

	// Load the .env file
	if err := godotenv.Load(*filename); err != nil {
		log.Fatalf("Error loading env file: %v", err)
	}

	// Parse the configuration file
	// TODO: get the value from env and store it in Config struct don't forget about the default value
}
