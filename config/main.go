package config

type AppConfig struct {
	Database	Database
}

func ReadEnvVar() (*AppConfig, error) {
	conf := AppConfig{}
	conf.Database = DatabaseEnv()
	return &conf, nil
}
