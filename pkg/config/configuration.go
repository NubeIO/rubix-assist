package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path"
)

var Config *Configuration
var RootCmd *cobra.Command

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Path     PathConfiguration
}

// Setup initialize configuration
func Setup(rootCmd_ *cobra.Command) error {
	RootCmd = rootCmd_
	configuration := &Configuration{}

	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.name", "data.db")

	Config = configuration
	return nil
}

func (conf *Configuration) Prod() bool {
	return RootCmd.PersistentFlags().Lookup("prod").Value.String() == "true"
}

func (conf *Configuration) GetPort() string {
	return RootCmd.PersistentFlags().Lookup("port").Value.String()
}

func (conf *Configuration) GetAbsDataDir() string {
	return path.Join(conf.getGlobalDir(), conf.getDataDir())
}

func (conf *Configuration) GetAbsConfigDir() string {
	return path.Join(conf.getGlobalDir(), conf.getConfigDir())
}

func (conf *Configuration) getGlobalDir() string {
	rootDir := RootCmd.PersistentFlags().Lookup("root-dir").Value.String()
	appDir := RootCmd.PersistentFlags().Lookup("app-dir").Value.String()
	return path.Join(rootDir, appDir)
}

func (conf *Configuration) getDataDir() string {
	return RootCmd.PersistentFlags().Lookup("data-dir").Value.String()
}

func (conf *Configuration) getConfigDir() string {
	return RootCmd.PersistentFlags().Lookup("config-dir").Value.String()
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
