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
	var config tomlConfig
	if _, err := toml.DecodeFile("./local.config.toml", &config); err != nil {
		Println(err)
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

