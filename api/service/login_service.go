package service

import (
    "carver_golang_blog/api/model"
    "carver_golang_blog/config"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    //."fmt"
)

/**
 *@desc 返回接口信息
 *@author Carver
 */
type returnJson struct {
    Code int
    Msg  string
    Data returnData
}

type returnData struct {
    UserId   int
    UserName string
}

/**
 * desc 处理用户登录信息
 * author Carver
 */
func GetUserInfo(formData *gin.Context) []returnJson {

    db := config.DB
    var result []returnJson
    var userInfo model.UserField

    //接收参数
    userName := formData.PostForm("userName")
    userPwd := formData.PostForm("userPwd")

    //验证用户是否存在
    row := db.QueryRow("select user_id,user_name,user_pwd from carver_user where user_name = ?", userName)

    err := row.Scan(&userInfo.UserId, &userInfo.UserName, &userInfo.UserPwd)

    if err != nil {
        var jsonError returnJson
        jsonError.Code = 5002
        jsonError.Msg = "用户名不存在~"
        result = append(result, jsonError)

    } else {
        //验证用户密码
        match := PasswordVerify(userPwd, userInfo.UserPwd)
        if match == false {
            var jsonError returnJson
            jsonError.Code = 5002
            jsonError.Msg = "密码不正确~"
            result = append(result, jsonError)

        } else {
            var jsonSuccess returnJson
            jsonSuccess.Code = 200
            jsonSuccess.Msg = "登录成功!"
            //返回用户信息
            jsonSuccess.Data.UserId = userInfo.UserId
            jsonSuccess.Data.UserName = userInfo.UserName

            result = append(result, jsonSuccess)

        }
    }

    return result
}

/**
 *@desc 进行hash加密
 *@author Carver
 */
func PasswordHash(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

/**
 *@desc 进行hash解密
 *@author Carver
 */
func PasswordVerify(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
