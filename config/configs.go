package config

import (
	"database/sql"
	. "fmt"
	 _"github.com/go-sql-driver/mysql"
	"github.com/BurntSushi/toml"
	//"path/filepath"
)

type tomlConfig struct {
	Title   string
	Api  	apiInfo
	Db      database `toml:"database.main"`
}

type apiInfo struct {
	ListenPort  string
	Debug  bool
	LogLevel  string
}

type database struct {
	DSN  string
	Active  string
}


func init() {

    pwd, err_pwd := os.Getwd()
        if err_pwd != nil {
             Println(err_pwd)
             os.Exit(1)
        }

    env_url := pwd + "/config/local.config.toml"

    var config tomlConfig
    _, err_config := toml.DecodeFile(env_url, &config)

    if err_config != nil {
    	Println(err_config)
    	return
    }



	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	db, _ := sql.Open("mysql", "用户名:密码@(localhost)/数据库") // 设置连接数据库的参数
	defer db.Close()                                            //关闭数据库
	err := db.Ping()                                            //连接数据库
	if err != nil {
		Println("数据库连接失败")
		return
	} else {
		Println("数据库连接成功")
	}


}

