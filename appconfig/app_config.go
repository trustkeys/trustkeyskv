package appconfig

import (
	"fmt"
	"log"
	"os"
	"strconv" )

var (
	// Merged Config pointer
	AppConfig *MergedConfig

	// Base Setting
	RunningHost   string
	ListenPort    string
	ServeDocument bool
	DocPath       string

	LogName string
	LogPath string
	KeepLog int

	MaxProcs      int
	MaxIdle       int
	MinIdle       int

	SystemOperatorID int64
	// Data Setting
	DefaultPassword      string
	DefaultAdminPassword string

	HTTPProxy           string
	BIGSETKV_HOST		string 
	BIGSETKV_PORT		int
	BIGSETKV_PATH 		string //Zookeeper path or etcd path
)


func InitConfig() {
	fmt.Println("[TrustKeyKV] Start loading host - port configuration.")

	var err error
	AppConfig, err = LoadConfig(os.Args[0] + ".conf")
	if err != nil || AppConfig == nil {
		log.Fatalln("[TrustKeyKV] Failed to load host - port configuration")
	}

	// Load base setting
	RunningHost = AppConfig.StringDefault("base_setting", "running_host", "0.0.0.0")
	ListenPort = AppConfig.StringDefault("base_setting", "listen_port", "3000")

	ServeDocument = AppConfig.BoolDefault("base_setting", "serve_doc", false)
	DocPath = AppConfig.StringDefault("base_setting", "doc_path", "/")

	LogName = AppConfig.StringDefault("base_setting", "log_name", "simplekvserver.log")
	LogPath = AppConfig.StringDefault("base_setting", "log_path", os.Args[0]+"/logs")

	KeepLog, err = strconv.Atoi(AppConfig.StringDefault("base_setting", "keep_log", "365"))
	if err != nil {
		MaxProcs = 200
	}

	MaxProcs, err = strconv.Atoi(AppConfig.StringDefault("base_setting", "maxprocs", "1"))
	if err != nil {
		MaxProcs = 1
	}


	MaxIdle, err = strconv.Atoi(AppConfig.StringDefault("base_setting", "max_idle", "8"))
	if err != nil {
		MaxIdle = 8
	}

	MinIdle, err = strconv.Atoi(AppConfig.StringDefault("base_setting", "min_idle", "0"))
	if err != nil {
		MinIdle = 0
	}

	HTTPProxy = AppConfig.StringDefault("base_setting", "proxy", "")

	BIGSETKV_HOST = AppConfig.StringDefault("host_port", "bigsetkv_host", "")
	BIGSETKV_PORT = AppConfig.IntDefault("host_port", "bigsetkv_port", 0)
	BIGSETKV_PATH = AppConfig.StringDefault("zookeeper", "bigsetkv_path", "/openstars/bigsetkv")
}
