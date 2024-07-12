package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	ProductionEnv = "production"
	HttpVersionV1 = "http1.1"
	HttpVersionV2 = "http2"
)

type Schema struct {

	// proxy
	ProxyHttpVersion  string `env:"proxy_http_version" envDefault:"http1.1"`
	ProxyEnableTLS    bool   `env:"proxy_enable_tls" envDefault:"false"`
	ProxyCertFile     string `env:"proxy_cert_file"`
	ProxyKeyFile      string `env:"proxy_key_file"`
	ProxyReadTimeout  int    `env:"proxy_read_timeout" envDefault:"10"`
	ProxyWriteTimeout int    `env:"proxy_write_timeout" envDefault:"10"`
	ProxyIdleTimeout  int    `env:"proxy_idle_timeout" envDefault:"90"`
	ProxyConcurrency  int    `env:"proxy_concurrency" envDefault:"100"`
	ProxyHttpPort     int    `env:"proxy_http_port" envDefault:"7777"`

	// redis
	RedisURI                string `env:"redis_uri" envDefault:"localhost:6379"`
	RedisStreamGlobalPrefix string `env:"redis_stream_global_prefix" envDefault:"tyrscale:stream"`
	RedisDBGlobalPrefix     string `env:"redis_db_global_prefix" envDefault:"tyrscale:db"`
	RedisPassword           string `env:"redis_password"`
	RedisDB                 int    `env:"redis_db" envDefault:"0"`

	// tyrscaleApi
	TyrscaleApiUrl string `env:"tyrscale_api_url" envDefault:"http://localhost:8888/api/v1"`

	// logging
	LogLevel string `env:"log_level" envDefault:"debug"`

	// Global
	ServerName  string `env:"server_name"`
	AppVersion  string `env:"app_version"`
	Environment string `env:"environment" envDefault:"production"`

	// otlp
	OtlpEnabled  bool   `env:"otlp_enabled" envDefault:"false"`
	OtlpEndpoint string `env:"otlp_endpoint" envDefault:"http://localhost:4318/v1/traces"`

	// plugin
	PluginPath string `env:"plugin_path" envDefault:"./plugins"`
}

var (
	cfg     Schema
	cfgPath = flag.String("config", "./config/config.yaml", "path to the configuration file")
)

func LoadConfig() (*Schema, error) {
	flag.Parse()

	err := godotenv.Load(*cfgPath)
	if err != nil {
		log.Printf("Error on load configuration file, error: %v", err)
		//return nil, err
	}

	if err := env.Parse(&cfg); err != nil {
		//log.Fatalf("Error on parsing configuration file, error: %v", err)
		return nil, err
	}

	return &cfg, nil
}

func GetConfig() *Schema {
	return &cfg
}
