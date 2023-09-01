package env

import (
	"fmt"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var confName string = ".bot"
var confType string = "env"
var envPath string = "./"
var c *viper.Viper
var mtx sync.Mutex

func init() {
	loadEnv()
}

func SetEnv(key, value string) {
	defer mtx.Unlock()
	mtx.Lock()
	c.Set(key, value)
	err := c.WriteConfig()
	if err != nil {
		logrus.Warn("set os error", err)
	}
	os.Setenv(key, value)

}

func loadEnv() {
	defer mtx.Unlock()
	mtx.Lock()
	exists := pathExists(envPath)
	if !exists {
		os.Mkdir(envPath, 0775)
	}
	filePath := fmt.Sprintf("%v%v.%v", envPath, confName, confType)
	exists = pathExists(filePath)
	if !exists {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()
	}
	c = viper.New()
	c.AddConfigPath(envPath)
	c.SetConfigName(confName)
	c.SetConfigType(confType)
	if err := c.ReadInConfig(); err != nil {
		logrus.Warn("init env err", err)
	}
	keys := c.AllKeys()
	if len(keys) > 0 {
		for _, key := range keys {
			existsValue := os.Getenv(key)
			if existsValue == "" {
				os.Setenv(key, c.GetString(key))
			}
		}
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
