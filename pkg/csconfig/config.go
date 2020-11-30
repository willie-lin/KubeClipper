package csconfig

//
type GlobalConfig struct {
	Self     *string      `yaml:"-"`
	DbConfig *DatabaseCfg `yaml:"db_config,omitempty"`
}

//func (c *GlobalConfig)  Dump() error {
//
//	out, err := yaml.Marshal(c)
//	if err != nil {
//		return errors.Wrap(err, "failed marshaling config")
//		}
//	fmt.Printf("%s", string(out))
//	return nil
//}
//
//func (c *GlobalConfig) LoadConfigurationFile(path string) error  {
//
//	fcontent, err := ioutil.ReadFile(path)
//	if err != nil {
//		return errors.Wrap(err, "failed to read config file")
//	}
//	err = yaml.UnmarshalStrict(fcontent, c)
//	if err != nil {
//		return errors.Wrap(err, "failed unmarshaling config")
//	}
//	path, err = filepath.Abs(path)
//	if err != nil {
//		return errors.Wrap(err, "failed to load absolute path")
//	}
//	c.Self = &path
//	if err := c.LoadConfiguration(); err != nil {
//		return errors.Wrap(err, "failed to load sub configurations")
//	}
//	return nil
//}
//
//func (c *GlobalConfig) LoadConfiguration() error {
//
//
//}

func NewConfig() *GlobalConfig {
	cfg := GlobalConfig{}
	return &cfg
}

func NewDefaultConfig() *GlobalConfig {
	//logLevel  := log.InfoLevel

	dbConfig := DatabaseCfg{
		User:     "root",
		Password: "root1234",
		Host:     "127.0.0.1",
		Port:     3306,
		DbName:   "KubeClipper",
	}
	globalCfg := GlobalConfig{
		DbConfig: &dbConfig,
	}

	return &globalCfg

}
