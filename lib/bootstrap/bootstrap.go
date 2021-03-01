package bootstrap

import (
	"os"

	"github.com/iconmobile-dev/backend-coding-challenge/config"
	"github.com/iconmobile-dev/go-core/logger"
)

// LoggerAndConfig is returning logger & config which are
// required for bootstrapping a service server
// is using CONFIG_FILE env var and if not set uses
// cfgFilePath. If cfgFilePath is not set then it tries to find the config
func LoggerAndConfig(serverName string, test bool) (logger.Logger, config.Config) {
	// init logger
	log := logger.Logger{MinLevel: "verbose"}

	// load config
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		log.Error("CONFIG_FILE env var missing\nuse: export CONFIG_FILE=<config_file_path>")
		os.Exit(1)
	}
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Error(err, "config load")
		os.Exit(1)
	}

	// set service name
	log.MinLevel = cfg.Logging.MinLevel
	log.TimeFormat = cfg.Logging.TimeFormat
	log.UseColor = cfg.Logging.UseColor
	log.ReportCaller = cfg.Logging.ReportCaller
	cfg.Server.Name = serverName

	if test {
		log.UseColor = false
		cfg.Server.Name += "_test"
	}

	if cfg.Server.Env == "prod" {
		log.UseJSON = true
	}

	return log, *cfg
}