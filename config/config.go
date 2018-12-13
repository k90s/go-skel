package config

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/pflag"

	"github.com/spf13/viper"
)

var (
	C config
	// Root 项目跟目录
	Root string
)

func init() {
	Root = os.Getenv("GOPATH") + "/src/github.com/kai-ding/go-skel"
	pflag.String("config", filepath.Join(Root, "config/config.yml"), "config file path")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	configFile := viper.GetString("config")

	if config := os.Getenv("CONFIG"); config != "" {
		configFile = config
	}

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic("failed to read config file: " + err.Error())
	}
	if err := viper.Unmarshal(&C); err != nil {
		panic("failed to unmarshal config: " + err.Error())
	}

	if port := os.Getenv("port"); port != "" {
		C.Server.Port, _ = strconv.Atoi(port)
	}
}

type (
	config struct {
		Server   Server
		Postgres Postgres
		JWT      JWT
	}

	Server struct {
		Port  int
		Debug bool
	}

	Postgres struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		LogMode  bool
		MaxOpen  int
		MaxIdle  int
	}

	JWT struct {
		Secret    string
		TTL       int64
		WhiteList []string
	}
)
