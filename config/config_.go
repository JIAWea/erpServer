package config

import (
	"flag"
	"fmt"
	"github.com/JIAWea/erpServer/pkg/utils"
	log "github.com/ml444/glog"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strconv"
)

const (
	EnvAuthJwtPem      = "AUTH_JWT_PEM"
	EnvKeyDebug        = "SERVICE_DEBUG"
	EnvKeyServiceDbDSN = "SERVICE_DB_DSN"
	EnvKeyASSET        = "SERVICE_ASSET"
)

// Config contains the fields for running a server
type Config struct {
	Debug      bool   `yaml:"debug"`
	EnableHTTP bool   `yaml:"enable_http"`
	EnableGRPC bool   `yaml:"enable_grpc"`
	HTTPAddr   string `yaml:"http_addr"`
	PprofAddr  string `yaml:"pprof_addr"`
	GRPCAddr   string `yaml:"grpc_addr"`

	JwtPem string `yaml:"jwt_pem"`

	Trans struct {
		PubKey string `yaml:"public_key"`
		PriKey string `yaml:"private_key"`
	} `yaml:"transfer_encrypt"`

	DbDSN    string `yaml:"db_dsn"`
	RedisDSN string `yaml:"redis_dsn"`
	AssetDir string `yaml:"asset_dir"`
}

var configFile string
var DefaultConfig Config

func init() {
	// NOTE: Flags have priority over Env vars.
	flag.StringVar(&configFile, "c", "./config.yaml", "config file")
	flag.BoolVar(&DefaultConfig.Debug, "debug", false, "enable APIs for pprof")
	flag.BoolVar(&DefaultConfig.EnableHTTP, "enable_http", true, "enable APIs for http")
	flag.BoolVar(&DefaultConfig.EnableGRPC, "enable_grpc", true, "enable APIs for grpc")
	flag.StringVar(&DefaultConfig.PprofAddr, "pprof_addr", ":5060", "Debug and metrics listen address")
	flag.StringVar(&DefaultConfig.HTTPAddr, "http_addr", ":5050", "HTTP listen address")
	flag.StringVar(&DefaultConfig.GRPCAddr, "grpc_addr", ":5040", "gRPC (HTTP) listen address")

}

func Init() {
	file, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Println("fail to read file:", err)
		os.Exit(-1)
	}

	err = yaml.Unmarshal(file, &DefaultConfig)
	if err != nil {
		fmt.Println("fail to yaml unmarshal:", err)
		os.Exit(-1)
	}

	parseEnvConfig()

	_ = os.Setenv(EnvAuthJwtPem, DefaultConfig.JwtPem)

	assDir := []string{"detail", "import"}
	for _, v := range assDir {
		path := filepath.Join(DefaultConfig.AssetDir, v)
		exist, _ := utils.IsPathExist(path)
		if !exist {
			_ = os.MkdirAll(path, 0777)
		}
	}
}

func parseEnvConfig() {
	// Use environment variables, if set.
	if enable := os.Getenv(EnvKeyDebug); enable != "" {
		debug, err := strconv.ParseBool(enable)
		if err != nil {
			log.Errorf("err: %v", err)
		} else {
			DefaultConfig.Debug = debug
		}
	}
	if enable := os.Getenv("ENABLE_HTTP"); enable != "" {
		e, err := strconv.ParseBool(enable)
		if err != nil {
			log.Errorf("err: %v", err)
		} else {
			DefaultConfig.EnableHTTP = e
		}
	}
	if enable := os.Getenv("ENABLE_GRPC"); enable != "" {
		e, err := strconv.ParseBool(enable)
		if err != nil {
			log.Errorf("err: %v", err)
		} else {
			DefaultConfig.EnableGRPC = e
		}
	}
	if addr := os.Getenv("PPROF_ADDR"); addr != "" {
		DefaultConfig.PprofAddr = addr
	}
	if addr := os.Getenv("HTTP_ADDR"); addr != "" {
		DefaultConfig.HTTPAddr = addr
	}
	if addr := os.Getenv("GRPC_ADDR"); addr != "" {
		DefaultConfig.GRPCAddr = addr
	}
	if DefaultConfig.DbDSN == "" {
		DefaultConfig.DbDSN = os.Getenv(EnvKeyServiceDbDSN)
		if DefaultConfig.Debug {
			log.Info("DB:", DefaultConfig.DbDSN)
		}
	}
}

func GetConfig() *Config {
	return &DefaultConfig
}
