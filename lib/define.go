package lib

import (
    "os"
    "github.com/gin-gonic/gin"
)

//框架的一些全局定义

const (
    C_Mysql = "mysql.json"
    C_Memcached = "memcached.json"
    C_Redis = "redies.json"
    C_Oss = "oss.json"
    C_Ots = "ots.json"
    C_Mns = "mns.json"

    DedfaultEnv = "dev"
)

//全局变量
var ENV_GIN = map[string]string{
    "dev":gin.DebugMode,
    "test":gin.TestMode,
    "staging":gin.ReleaseMode,
    "pro":gin.ReleaseMode,
}

type Define struct {
    DefaultEnv          string
    Env                 string
    EnvPath             string
    EnvConfigPath       string

    ConfigPath          string

    MysqlConfigPath     string
    MemcachedConfigPath string
    RedisConfigPath     string
    OssConfigPath       string
    OtsConfigPath       string
    MnsConfigPath       string
}

var TotalDefine = &Define{DefaultEnv: DedfaultEnv}

func (d *Define)InitConfig(envPath string) {
    d.setEnv(envPath)
    d.ConfigPath = Dir(envPath)
    d.EnvConfigPath = d.ConfigPath + string(os.PathSeparator) + d.Env
    d.EnvPath = Dir(envPath)

    d.MysqlConfigPath = d.EnvConfigPath + string(os.PathSeparator) + C_Mysql
    d.MemcachedConfigPath = d.EnvConfigPath + string(os.PathSeparator) + C_Memcached
    d.RedisConfigPath = d.EnvConfigPath + string(os.PathSeparator) + C_Redis
    d.MnsConfigPath = d.EnvConfigPath + string(os.PathSeparator) + C_Mns
    d.OssConfigPath = d.EnvConfigPath + string(os.PathSeparator) + C_Oss
    d.OtsConfigPath = d.EnvConfigPath + string(os.PathSeparator) + C_Ots

}

func (d *Define)setEnv(envPath string) {
    env := ReadFile(envPath)
    if env == "" {
        env = d.DefaultEnv
    } else {
        d.Env = env
    }
}

