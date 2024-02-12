package config

type Configuration struct {
	RawURL           string `mapstructure:"RAW_URL"`
	LogType          string `mapstructure:"LOG_TYPE"`
	LogOutput        string `mapstructure:"LOG_OUTPUT"`
	LogPrefix        string `mapstructure:"LOG_PREFIX"`
	LogFileName      string `mapstructure:"LOG_FILE_NAME"`
	LogFileMaxSize   int    `mapstructure:"LOG_FILE_MAX_SIZE"`
	LogFileMaxBackup int    `mapstructure:"LOG_FILE_MAX_BACKUP"`
	LogFileMaxAge    int    `mapstructure:"LOG_FILE_MAX_AGE"`
	LogFileCompress  bool   `mapstructure:"LOG_FILE_COMPRESS"`
	URI              string `mapstructure:"URI"`
	TestnetURI       string `mapstructure:"TESTNET_URI"`
	UseTestnet       bool   `mapstructure:"USE_TESTNET"`
}

var DEFAULTS = map[string]interface{}{
	"RAW_URL":             "https://cloudflare-eth.com",
	"LOG_TYPE":            "debug",
	"LOG_PREFIX":          "client_svc",
	"LOG_OUTPUT":          "stdout",
	"LOG_FILE_NAME":       "./client.log",
	"LOG_FILE_MAX_SIZE":   50,
	"LOG_FILE_MAX_BACKUP": 5,
	"LOG_FILE_MAX_AGE":    30,
	"LOG_FILE_COMPRESS":   true,
	"LOG_MAX_BACKUPS":     50,
	"LOG_MAX_AGE":         30,
	"URI":                 "https://api.binance.com",
	"TESTNET_URI":         "https://testnet.binance.vision",
	"USE_TESTNET":         false,
}
