package config

type ServerConfiguration struct {
	ListenAddr                 string `default:"0.0.0.0"`
	Port                       string
	Secret                     string
	AccessTokenExpireDuration  int     `default:"1"`
	RefreshTokenExpireDuration int     `default:"1"`
	LimitCountPerRequest       float64 `default:"0"`
}
