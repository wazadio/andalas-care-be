package env

import "github.com/joho/godotenv"

type envConfig struct {
	Filename string
}

func NewEnvConfig(filename string) *envConfig {
	return &envConfig{
		Filename: filename,
	}
}

func (e *envConfig) LoadEnv() error {
	var err error
	if e.Filename == "" {
		err = godotenv.Load()
	} else {
		err = godotenv.Load(e.Filename)
	}

	return err
}
