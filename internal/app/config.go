package app

type Config struct {
	RunPort  string `yaml:"run_port"`
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
}
