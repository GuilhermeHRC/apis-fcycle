package configs

import (
	"fmt"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`

	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`

	TokenAuth *jwtauth.JWTAuth `mapstructure:"-"`
}

var cfg *conf

func LoadConfig(path string) (*conf, error) {
	v := viper.New()

	v.SetConfigFile(path + "/.env")
	v.SetConfigType("env")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Nenhum .env encontrado, usando apenas variáveis de ambiente!")
	}

	var c conf
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("erro ao mapear configuração: %w", err)
	}

	// JWT Auth
	c.TokenAuth = jwtauth.New("HS256", []byte(c.JWTSecret), nil)

	cfg = &c
	return cfg, nil
}
