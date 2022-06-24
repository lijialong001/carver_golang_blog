package controller
import (
    "github.com/gin-gonic/gin"
    "carver_golang_blog/api/service"
    //"io/ioutil"
    "net/http"
    ."fmt"
)


//用户登录
func Login(context *gin.Context) {
     userInfo := service.GetUserInfo()
     Prinln(userInfo)
    //context.JSON(http.StatusOK, gin.H{
    //    	"data": userInfo,
    //})

    //获取请求body
//    user_params, _ := ioutil.ReadAll(context.Request.Body)
//    user_info := string(user_params)

//    context.JSON(http.StatusOK, gin.H{
//	"message": user_info,
//    })
}


