package config

type AppConfig struct {
	App struct {
		Port           int
		Debug          bool
		Authentication struct {
			Username string
			Password string
		}
	}
	Config struct {
		Filename string
	}
}
