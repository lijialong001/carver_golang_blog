package service

import (
    "carver_golang_blog/config"
    "carver_golang_blog/api/model"
    ."fmt"
)




func GetUserInfo(){
    result := model.UserField
    db := config.DB
    rows, _ := db.Query("select * from test_city") //获取city表所有数据
    for rows.Next() {
       var f UserField
       rows.Scan(&f.Id ,&f.Name, &f.Pid)
       result = append(result,f)
    }

    return  result
}
