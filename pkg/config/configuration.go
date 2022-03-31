package config

import (
	"flag"
	"github.com/NubeDev/configor"
	"path"
)

var config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Path     PathConfiguration
	Prod     bool `default:"false"`
}

// Setup initialize configuration
func Setup() *Configuration {
	config = new(Configuration)
	config = config.Parse()
	err := configor.New(&configor.Config{EnvironmentPrefix: "ASSIST"}).Load(config, path.Join(config.GetAbsConfigDir(), "config.yml"))
	if err != nil {
		panic(err)
	}
	return config
}

func (conf *Configuration) Parse() *Configuration {
	port := flag.String("p", "8080", "Port")
	globalDir := flag.String("g", "./", "Global Directory")
	dataDir := flag.String("d", "data", "Data Directory")
	configDir := flag.String("c", "config", "Config Directory")
	prod := flag.Bool("prod", false, "Deployment Mode")
	flag.Parse()
	conf.Server.Port = *port
	conf.Path.GlobalDir = *globalDir
	conf.Path.DataDir = *dataDir
	conf.Path.ConfigDir = *configDir
	conf.Prod = *prod
	return conf
}

func (conf *Configuration) GetAbsDataDir() string {
	return path.Join(conf.Path.GlobalDir, conf.Path.DataDir)
}

func (conf *Configuration) GetAbsConfigDir() string {
	return path.Join(conf.Path.GlobalDir, conf.Path.ConfigDir)
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return config
}
