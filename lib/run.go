package lib

import (
    "log"
    "os"
    "os/signal"
    "syscall"
    "github.com/gin-gonic/gin"
    "time"
    "flag"
)

func Run(router *gin.Engine, port string) {

    if err := router.Run(port); err != nil {
        time.Sleep(time.Second)
        log.Fatal("service fail")
        return
    }

    c := make(chan os.Signal)
    signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)
    for s := range c {
        switch s {
        case syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT:
            log.Println("service end receive signal %v", s)
            return
        default:
            log.Println("other")
        }
    }
    return
}

func Start() {
    log.Println("####service--start####")
    var envp string
    flag.StringVar(&envp, "envp", "", "please input --envp=`you config path`")
    flag.Parse()
    if envp == "" {
        panic("name is not set :add  --envp=`you config path`  after go run main ")
    }
    //初始化配置
    TotalDefine.InitConfig(envp)
}
