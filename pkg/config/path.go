package config

type PathConfiguration struct {
	Local      string
	ToPath     string `default:"data"`
	FromPath   string `default:"data"`
	UnZipPath  string `default:"data"`
	FlowPlugin string `default:"/data/flow-framework/data/plugins"`
	GlobalDir  string `default:"./"`
	ConfigDir  string `default:"config"`
	DataDir    string `default:"data"`
}
