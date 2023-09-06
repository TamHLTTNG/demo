package config

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Server struct {
		Address          string `mapstructure:"SERVER_ADDRESS"`
		OpenAM           string `mapstructure:"SERVER_OPENAM"`
		KeyCloakBaseURL  string `mapstructure:"SERVER_KEYCLOAK"`
		KeyCloakRealm    string `mapstructure:"SERVER_KEYCLOAK_REALM"`
		KeyCloakClientID string `mapstructure:"SERVER_KEYCLOAK_CLIENTID"`
		KeyCloakSecret   string `mapstructure:"SERVER_KEYCLOAK_SECRET"`
	}

	PGADMIN struct {
		Email    string `mapstructure:"PGADMIN_DEFAULT_EMAIL"`
		Password string `mapstructure:"PGADMIN_DEFAULT_PASSWORD"`
	}

	DB struct {
		Host     string `mapstructure:"DB_HOST"`
		Driver   string `mapstructure:"DB_DRIVER"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
		Name     string `mapstructure:"DB_NAME"`
		Port     string `mapstructure:"DB_PORT"`
	}
}

var C config

func ReadConfig() {
	Config := &C

	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config.Server); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&Config.DB); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&Config.PGADMIN); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}
