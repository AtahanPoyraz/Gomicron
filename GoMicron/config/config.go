package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

//---[ SERVER SETTINGS ]---------------------------------------------------------------------------------------------------------------------------//

type ServerConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	IdleTimeout  int 	`yaml:"idletimeout"`
	ReadTimeout  int 	`yaml:"readtimeout"`
	WriteTimeout int 	`yaml:"writetimeout"`
}

//---[ CORS SETTINGS ]------------------------------------------------------------------------------------------------------------------------------//

type CORSConfig struct {
	AllowedOrigins []string `yaml:"allowed_origins"`
	AllowedMethods []string `yaml:"allowed_methods"`
	AllowedHeaders []string `yaml:"allowed_headers"`
}

//---[ SERVICE SETTINGS ]--------------------------------------------------------------------------------------------------------------------------//

type MySQLConfig struct {
	DBHost         string `yaml:"db_host"`
	DBPort         int    `yaml:"db_port"`
	DBUser         string `yaml:"db_user"`
	DBPass         string `yaml:"db_pass"`
	DBName         string `yaml:"db_name"`
	DBTimeout      int    `yaml:"db_timeout"`
	DBSSLMode      string `yaml:"db_sslmode"`
	DBNET		   string `yaml:"db_net"`
	DBSrvname      string `yaml:"db_srvname"`
}

type PostgresConfig struct {
	DBHost         string `yaml:"db_host"`
	DBPort         int    `yaml:"db_port"`
	DBUser         string `yaml:"db_user"`
	DBPass         string `yaml:"db_pass"`
	DBName         string `yaml:"db_name"`
	DBTimeout      int    `yaml:"db_timeout"`
	DBSSLMode      string `yaml:"db_sslmode"`
	DBNET		   string `yaml:"db_net"`
	DBSrvname      string `yaml:"db_srvname"`
}

type Services struct {
	Postgres PostgresConfig `yaml:"postgres"`
	MySql    MySQLConfig    `yaml:"mysql"` 
}

//---[ GENERAL SETTINGS STRUCT ]-------------------------------------------------------------------------------------------------------------------//

type Config struct {
	Server   ServerConfig `yaml:"server"`
	CORS     CORSConfig   `yaml:"cors"`
	Services Services    `yaml:"services"`
}
//---[ CONFIG READER ]----------------------------------------------------------------------------------------------------------------------------//

func ReadConfigFromFile(filePath string) (Config, error) {
	var config Config

	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}