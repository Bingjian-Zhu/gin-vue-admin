package setting

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

//Config 定义配置
var (
	Config *Conf
)

//Conf 配置结构体
type Conf struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
	RabbitMQ RabbitMQ `yaml:"rabbitmq"`
	APP      APP      `yaml:"app"`
}

//Server HTTP服务配置结构体
type Server struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read-timeout"`
	WriteTimeout time.Duration `yaml:"write-timeout"`
}

//APP 程序配置
type APP struct {
	RunMode     string `yaml:"run-mode"`
	Pagesize    int    `yaml:"pagesize"`
	IdentityKey string `yaml:"identity-key"`
	LogPath     string `yaml:"log-path"`
}

//Database 数据库配置结构体
type Database struct {
	Type        string `yaml:"type"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"table-prefix"`
}

//Redis redis配置结构体
type Redis struct {
	Addr        string        `yaml:"addr"`
	Pass        string        `yaml:"pass"`
	DB          int           `yaml:"db"`
	Timeout     time.Duration `yaml:"timeout"`
	ExpiredTime int           `yaml:"expired-time"`
}

//RabbitMQ RabbitMQ配置结构体
type RabbitMQ struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

//init 初始化函数
func init() {
	Config = getConf()
	log.Println("[Setting] Config init success")
}

//getConf 读取配置文件
func getConf() *Conf {
	var c *Conf
	file, err := ioutil.ReadFile("../config/config.yml")
	if err != nil {
		log.Println("[Setting] config error: ", err)
	}
	err = yaml.UnmarshalStrict(file, &c)
	if err != nil {
		log.Println("[Setting] yaml unmarshal error: ", err)
	}
	return c
}
