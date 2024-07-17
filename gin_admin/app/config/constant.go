package config

const (
	ConfigEnv         = "CONFIG"
	ConfigDefaultFile = "config.yaml"
	ConfigTestFile    = "config.test.yaml"
	ConfigDebugFile   = "config.debug.yaml"
	ConfigReleaseFile = "config.release.yaml"
)
type JWT struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Expire int64  `mapstructure:"expire" json:"expire" yaml:"expire"`
}

type Config struct {
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	GeneralDB
}