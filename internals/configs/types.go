package configs

type (
	Config struct {
		Service  Service  `yaml:"service"`
		Database Database `yaml:"database"`
	}

	Service struct {
		Port string `mapstructure:"port"`
	}

	Database struct {
		DataSourceName string `mapstructure:"dataSourceName"`
	}
)
