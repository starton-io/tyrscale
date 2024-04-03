package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	ProductionEnv = "production"
)

type Schema struct {
	Environment string `env:"environment" envDefault:"production"`
	HttpPort    int    `env:"http_port" envDefault:"8888"`

	//GrpcPort          int    `env:"grpc_port"`
	RedisURI                string `env:"redis_uri" envDefault:"localhost:6379"`
	RedisStreamGlobalPrefix string `env:"redis_stream_global_prefix" envDefault:"tyrscale:stream"`
	RedisDBGlobalPrefix     string `env:"redis_db_global_prefix" envDefault:"tyrscale:db"`
	LogLevel                string `env:"log_level" envDefault:"debug"`
	RedisPassword           string `env:"redis_password"`
	RedisDB                 int    `env:"redis_db" envDefault:"0"`
	ReadTimeout             int    `env:"read_timeout" envDefault:"10"`
	WriteTimeout            int    `env:"write_timeout" envDefault:"10"`
	ServerName              string `env:"server_name"`
	OtlpEnabled             bool   `env:"otlp_enabled" envDefault:"false"`
	OtlpEndpoint            string `env:"otlp_endpoint" envDefault:"http://localhost:4318/v1/traces"`
	AppVersion              string `env:"app_version"`
}

var (
	cfg     Schema
	cfgPath = flag.String("config", "./config/config.yaml", "path to the configuration file")
)

func LoadConfig() (*Schema, error) {
	flag.Parse()

	err := godotenv.Load(*cfgPath)
	if err != nil {
		log.Printf("Error on loading config file, error: %v", err)
	}

	if err := env.Parse(&cfg); err != nil {
		//log.Fatalf("Error on parsing configuration file, error: %v", err)
		return nil, err
	}

	return &cfg, nil
}

func GetConfig() *Schema {
	fmt.Println(&cfg)
	return &cfg
}
