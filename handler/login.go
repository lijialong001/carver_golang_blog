package handler
import (
"github.com/gin-gonic/gin"
"io/ioutil"
"net/http"
//"fmt"
)

//用户登录
func Login(context *gin.Context) {
    //获取请求body
    user_params, _ := ioutil.ReadAll(context.Request.Body)
    user_info := string(user_params)
    context.JSON(http.StatusOK, gin.H{
	"message": user_info,
    })
}
