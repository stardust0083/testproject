package utils

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	//使用了beego框架的配置文件读取模块
)

var (
	G_server_name       string //项目名称
	G_server_addr       string //服务器ip地址
	G_server_port       string //服务器端口
	G_redis_addr        string //redis ip地址
	G_redis_port        string //redis port端口
	G_redis_dbnum       string //redis db 编号
	G_mysql_addr        string //mysql ip 地址
	G_mysql_port        string //mysql 端口
	G_mysql_dbname      string //mysql db name
	G_fastdfs_port      string //fastdfs 端口
	G_fastdfs_addr      string //fastdfs ip
	G_redis_maxidel     int
	G_redis_maxactive   int
	G_redis_idletimeout time.Duration
)

func InitConfig() {
	//从配置文件读取配置信息
	viper.SetConfigName("app")
	viper.SetConfigType("toml")                                       // 查找配置文件所在路径
	viper.AddConfigPath("/home/hsf/Desktop/project/testproject/conf") // 还可以在工作目录中搜索配置文件
	err := viper.ReadInConfig()                                       // 搜索并读取配置文件
	if err != nil {                                                   // 处理错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	G_server_name = viper.GetString("appname")
	G_server_addr = viper.GetString("httpaddr")
	G_server_port = viper.GetString("httpport")
	G_redis_addr = viper.GetString("redisaddr")
	G_redis_port = viper.GetString("redisport")
	G_redis_dbnum = viper.GetString("redisdbnum")
	G_mysql_addr = viper.GetString("mysqladdr")
	G_mysql_port = viper.GetString("mysqlport")
	G_mysql_dbname = viper.GetString("mysqldbname")
	G_fastdfs_port = viper.GetString("fastdfsport")
	G_fastdfs_addr = viper.GetString("fastdfsaddr")
	G_redis_maxidel = 10
	G_redis_maxactive = 0
	G_redis_idletimeout = time.Duration(1000)
	return
}

func init() {
	InitConfig()
}
