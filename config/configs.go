package config

import (
    "database/sql"
    "github.com/BurntSushi/toml"
    _ "github.com/go-sql-driver/mysql"
    "os"
    ."fmt"
)

type tomlConfig struct {
    Title   string
    Admin   adminInfo
    Db      database `toml:"database"`
    Redis   redisInfo
    Servers map[string]server
    Other   other
}

type adminInfo struct {
    Name     string
    Port     int `toml:"ListenPort"`
    Debug    bool
    LogLevel string
}

type database struct {
    UserName string `toml:"username"`
    Password string
    Database string
    Addr     string
    Ports    []int
    ConnMax  int `toml:"connection_max"`
    Enabled  bool
    Port     int `toml:"system_port"`
}

type redisInfo struct {
    Addr  string
    Auth  string
    Dial  string `toml:"DialTimeout"`
    Read  string `toml:"ReadTimeout"`
    Write string `toml:"WriteTimeout"`
}

type server struct {
    Db  string
    Pwd string
}

type other struct {
    Data [][]interface{}
    Type []string
}

var DB *sql.DB

/**
 *@desc 初始化数据库
 *@author Carver
 */
func init() {

    pwd, pwdError := os.Getwd()
    if pwdError != nil {
        os.Exit(1)
        panic(pwdError)
    }

    envUrl := pwd + "/config/local.config.toml"

    var config tomlConfig
    _, configError := toml.DecodeFile(envUrl, &config)

    if configError != nil {
        panic(configError)
    }

    //"用户名:密码@[连接方式](主机名:端口号)/数据库名"
    connStr := Sprintf("%s:%s@(%s)/%s", config.Db.UserName, config.Db.Password, config.Db.Addr, config.Db.Database)
    //设置连接数据库的参数
    db, _ := sql.Open("mysql", connStr)
    //连接数据库
    err := db.Ping()

    if err != nil {
        Println("数据库连接失败❌")
        return
    } else {

        Println("数据库连接成功✅")
    }

    DB = db

}
