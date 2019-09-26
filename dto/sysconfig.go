package dto

import (
	"github.com/BurntSushi/toml"
	"log"
)

var Config *SysConfig

type SysConfig struct {
	Service struct {
		Name        string //服务英文名
		DisplayName string //服务中文名
		Version     string //版本号
		Host        string
		Port        int
		AppId       string //项目编号
		EnvName     string //环境变量
	}

	//Consul服务地址
	Consul struct {
		Host string
		Port int
	}

	//数据库链接字符串
	Mysql struct {
		Host    string
		Port    int
		User    string
		Pwd     string
		Default string
	}

	//缓存链接字符串
	Cache struct {
		Host    string
		Port    int
		Pwd     string
		Default int
	}

	Rsa struct {
		PublicKey  string
		PrivateKey string
		Issuer     string
	}
	//数据库链接字符串
	MsSql struct {
		Host    string
		Port    int
		User    string
		Pwd     string
		Default string
	}
	ConfigDto struct {
		BrandUrl string
	}

	TaskConfig struct {
		ActivityConfig string
	}

	Grpc struct {
		Host       string
		CertFile   string
		ServerName string
	}
}

func (m *SysConfig) LoadConfig() (service SysConfig) {
	_, err := toml.DecodeFile("appsetting.toml", &m)
	if err != nil {
		log.Fatal(err)
	}
	Config = m
	return *m
}
