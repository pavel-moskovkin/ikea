package config

type Config struct {
	ApiListen string
	App       App `mapstructure:"app"`
	DB        DB  `mapstructure:"db"`
	LogLevel  string
}

func NewConfig() Config {
	return Config{
		App:       NewApp(),
		DB:        NewDB(),
		ApiListen: "8080",
		LogLevel:  "debug",
	}
}
