package config

import (
	"gonet/base"
	"gonet/base/system"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

var (
	env           string
	serviceConfig map[string]interface{}
	redisConfig   map[string]interface{}
	logConfig     map[string]interface{}
	mysqlConfig   map[string]interface{}
	mongoConfig   map[string]interface{}
	lock          sync.Mutex
)

func Init(_env string, _data interface{}) {
	env = _env
	load(_data)
}

func InitEnv(_env string) {
	env = _env
}

func ReadConf(path string, data interface{}) bool {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		//base.CLog.Fatalf("解析gonet.yaml读取错误: %v", err)
		return false
	}

	err = yaml.Unmarshal(content, data)
	if err != nil {
		//log.Fatalf("解析config.yaml出错: %v", err)
		return false
	}

	return true
}

func load(_data interface{}) {
	var serviceConfigPath = GetConfigPath("gonet.yaml")
	//var redisConfigPath = getConfigPath("redis.json")
	//var mysqlConfigPath = getConfigPath("mysql.json")
	//var mongoConfigPath = getConfigPath("mongo.json")
	//var logConfigPath = getConfigPath("log.json")

	lock.Lock()
	loadConfig(_data, serviceConfigPath)
	//loadConfig(&redisConfig, redisConfigPath)
	//loadConfig(&mysqlConfig, mysqlConfigPath)
	//loadConfig(&mongoConfig, mongoConfigPath)
	//loadConfig(&logConfig, logConfigPath)
	lock.Unlock()
}

func GetConfigPath(configFile string) string {
	return system.Root + "/config/" + env + "/" + configFile
}

/**
  返回表格数据文件地址
*/
func GetConfigTablePath(configDataFile string) string {
	return system.Root + "/config/" + env + "/" + "static/" + configDataFile
}

func loadConfig(data interface{}, configPath string) {
	base.ReadConf(configPath, data)
}

//
//func GetConnectorService(serviceId int) map[string]interface{} {
//	serviceData := serviceConfig["connector"].(map[string]interface{})
//	serverDatas := serviceData["services"].(map[string]interface{})
//	return serverDatas[NumToString(serviceId)].(map[string]interface{})
//}
//
//func GetConnectorServiceTslCrt() string {
//	serverData := serviceConfig["connector"].(map[string]interface{})
//	return serverData["tslCrt"].(string)
//}
//
//func GetConnectorServiceTslKey() string {
//	serverData := serviceConfig["connector"].(map[string]interface{})
//	return serverData["tslKey"].(string)
//}
//
//func GetApiService(serviceId int) map[string]interface{} {
//	serviceData := serviceConfig["api"].(map[string]interface{})
//	serverDatas := serviceData["services"].(map[string]interface{})
//	return serverDatas[NumToString(serviceId)].(map[string]interface{})
//}
//
//func GetApiServiceTslCrt() string {
//	serverData := serviceConfig["api"].(map[string]interface{})
//	return serverData["tslCrt"].(string)
//}
//
//func GetApiServiceTslKey() string {
//	serverData := serviceConfig["api"].(map[string]interface{})
//	return serverData["tslKey"].(string)
//}
//
//func GetRedisConfig() map[string]interface{} {
//	return redisConfig
//}
//
//func GetMysqlConfig() map[string]interface{} {
//	return mysqlConfig
//}
//
//func GetLogConfig() map[string]interface{} {
//	return logConfig
//}
//
//func GetMongoConfig() map[string]interface{} {
//	return mongoConfig
//}
