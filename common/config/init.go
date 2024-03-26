package config

import (
	"bytes"
	"code.freebrio.com/fb-go/lib/fbl"
	"embed"
	"fmt"
	"github.com/spf13/viper"

	_ "embed"
)

// Server 配置
type Server struct {
	AppName string `mapstructure:"appName" json:"address" yaml:"appName"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
}

// Consul consul配置
type Consul struct {
	Address    string `mapstructure:"address" json:"address" yaml:"address"`
	Datacenter string `mapstructure:"datacenter" json:"datacenter" yaml:"datacenter"`
	Token      string `mapstructure:"token" json:"token" yaml:"token"`
}

// Rpc rpc配置
type Rpc struct {
	ApiGolangAddress string `mapstructure:"api-golang" json:"api-golang" yaml:"api-golang"`
	ArenaAddress     string `mapstructure:"arena" json:"arena" yaml:"arena"`
	ApiServerAddress string `mapstructure:"api-server" json:"api-server" yaml:"api-server"`
}

// Mysql mysql配置
type Mysql struct {
	Dsn             string `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
	MaxIdleConn     int    `mapstructure:"maxIdleConn" json:"maxIdleConn" yaml:"maxIdleConn"`
	MaxOpenConn     int    `mapstructure:"maxOpenConn" json:"maxOpenConn" yaml:"maxOpenConn"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime" json:"connMaxLifetime" yaml:"connMaxLifetime"`
}

// Redis redis配置
type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码

	ReadTimeOut  int64 `mapstructure:"readTimeout" json:"readTimeout" yaml:"readTimeout"`
	WriteTimeout int64 `mapstructure:"writeTimeout" json:"writeTimeout" yaml:"writeTimeout"`
}

type Conf struct {
	Server *Server `mapstructure:"server" json:"server" yaml:"server"` // Server 配置

	Consul *Consul `mapstructure:"consul" json:"consul" yaml:"consul"` // Consul 配置

	Rpc *Rpc `mapstructure:"rpc" json:"rpc" yaml:"rpc"` // rpc 配置

	Mysql *Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"` // mysql 配置

	Redis *Redis `mapstructure:"redis" json:"redis" yaml:"redis"` // redis 配置
}

var Instance = initCf()

//go:embed *.yaml
var multiCf embed.FS

func initCf() *Conf {
	//获取环境变量
	env := fbl.FillValOfKey("API_ENV", "dev")
	fbl.Log().Sugar().Infof("start in %s", env)

	//读取配置文件
	data, err := multiCf.ReadFile("config." + env + ".yaml")
	if err != nil {
		fmt.Println(err)
	}

	//读取配置文件
	v := viper.New()
	v.SetConfigType("yaml")
	err = v.ReadConfig(bytes.NewBuffer(data))
	//err = v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	//解析配置文件
	var cf *Conf
	err = v.Unmarshal(&cf)
	if err != nil {
		panic(err)
	}

	//根据配置参数更新 consul token
	cf.Consul.Token = fbl.FillValOfKey("CONSUL_HTTP_TOKEN", cf.Consul.Token)
	return cf
}
