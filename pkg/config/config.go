package config

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/tkanos/gonfig"
)

type Config struct {
	BrokerHost     string
	BrokerPubTopic string
	BrokerSubTopic string
	BrokerPort     int
	DBHost         string
	DBName         string
	DBUser         string
	DBPwd          string
	DBPort         int
	DBSSL          string
	DBTmz          string
}

var Cfg = Config{}
var FilePath = flag.String("c", "./config.json", "config file path")

// func init() {
// 	InitConfig()
// }

func InitConfig() error {

	//reads the config file and loads the config struct
	if _, err := os.Stat(filepath.Join(".", "config.json")); !os.IsNotExist(err) {
		err = gonfig.GetConf(*FilePath, &Cfg)
		if err != nil {
			println("Error reading config file:", err)
		}
		println("Successfully read from config file: ", *FilePath)
	} else {
		println("Missing config file: ./config.json.")
	}

	return nil

}
