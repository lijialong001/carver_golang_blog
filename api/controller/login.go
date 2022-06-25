package controller

import (
    "carver_golang_blog/api/service"
    "github.com/gin-gonic/gin"
    //"io/ioutil"
    "net/http"
    //."fmt"
)

/**
 *@desc 用户登录
 *@author Carver
 */
func Login(context *gin.Context) {
    //处理用户信息
    userInfo := service.GetUserInfo(context)

    context.JSON(http.StatusOK, gin.H{
        "message": userInfo,
    })
}
