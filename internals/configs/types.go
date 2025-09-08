package configs

type (
	Config struct {
		Service  Service  `yaml:"service"`
		Database Database `yaml:"database"`
	}

	Service struct {
		Port string `mapstructure:"port"`
		SecretJWT string `mapstructure:"secretJWT"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}
)
