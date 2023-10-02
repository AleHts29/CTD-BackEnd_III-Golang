package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	PublicConfig  PublicConfig
	PrivateConfig PrivateConfig
}

type PublicConfig struct {
	PublicKey string
}

type PrivateConfig struct {
	SecretKey string
}

var (
	envs = map[string]PublicConfig{
		"local": {
			PublicKey: "localAdmin",
		},
		"dev": {
			PublicKey: "devAdmin",
		},
		"prod": {
			PublicKey: "prodAdmin",
		},
	}
)

func NewConfig(env string) (Config, error) {
	publicConfig, exists := envs[env]
	if !exists {
		return Config{}, errors.New("env doest not exists")
	}

	secretKey := os.Getenv("SECRET_KEY")
	fmt.Println("Secret_key::::", secretKey)
	if secretKey == "" {
		return Config{}, errors.New("secrect key does not exists in env")
	}

	return Config{
		PublicConfig: publicConfig,
		PrivateConfig: PrivateConfig{
			SecretKey: secretKey,
		},
	}, nil

}
