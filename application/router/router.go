package router

import (
    "github.com/gin-gonic/gin"
    "kobe/application/api"
    "kobe/lib"
    "log"
)

func InitRouter() *gin.Engine {
    gin.SetMode(lib.ENV_GIN[lib.TotalDefine.Env])//设置gin模式
    log.Println("####set--mode--", lib.ENV_GIN[lib.TotalDefine.Env], "####")
    router := gin.Default()
    router.GET("/", api.IndexApi)
    router.GET("/test", api.TestApi)

    return router
}