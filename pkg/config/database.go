package config

type DatabaseConfiguration struct {
	Driver   string `default:"sqlite"`
	Dbname   string `default:"test"`
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool `default:"true"`
}
