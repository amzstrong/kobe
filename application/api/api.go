package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "vendertest/application/model"
)

func IndexApi(c *gin.Context) {
    c.String(http.StatusOK, "It works")
}

func TestApi(c *gin.Context) {

    user := new(model.UserModel)
    res, _ := user.M().Fields("name").First()
    res2, _ := user.Con.Query("select * from user")
    c.JSON(http.StatusOK, gin.H{
        "msg": "test-ok",
        "res": res,
        "res2": res2,
    })

}


