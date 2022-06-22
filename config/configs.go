package config

import (
	"database/sql"
	. "fmt"
	 _"github.com/go-sql-driver/mysql"
	"github.com/BurntSushi/toml"
	//"path/filepath"
	"os"
)

type tomlConfig struct {
	Title       string
	Admin       adminInfo
	Db          database    `toml:"database"`
	Redis       redisInfo
	Servers     map[string]server
	Other       other
}

type adminInfo struct {
	Name        string
	Port        int     `toml:"ListenPort"`
	Debug       bool
	LogLevel    string

}

type database struct {
    UserName    string  `toml:"username"`
    Password    string
    Database    string
	Addr        string
	Ports       []int
	ConnMax     int     `toml:"connection_max"`
	Enabled     bool
	Port        int     `toml:"system_port"`
}

type redisInfo struct {
    Addr        string
    Auth        string
    Dial        string  `toml:"DialTimeout"`
    Read        string  `toml:"ReadTimeout"`
    Write       string  `toml:"WriteTimeout"`
}


type server struct {
	Db      string
	Pwd     string
}

type other struct {
	Data    [][]interface{}
	Type    []string
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
	db, _ := sql.Open("mysql", config.Db.UserName+":"+config.Db.Password+"@(localhost)/"+config.Db.Database) // 设置连接数据库的参数
	defer db.Close()                                            //关闭数据库
	err := db.Ping()                                            //连接数据库
	if err != nil {
		Println("数据库连接失败")
		return
	} else {
		Println("数据库连接成功")
	}


}

