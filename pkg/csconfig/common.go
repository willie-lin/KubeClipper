package csconfig

type CommonCfg struct {
	Daemonize  bool
	PidDir     string `yaml:"pid_dir"`
	LogMedia   string `yaml:"log_media"`
	LogDir     string `yaml:"log_dir, omitempty"`
	logLevel   string `yaml:"log_level"`
	WorkingDir string `yaml:"working_dir, omitempty"`
}
