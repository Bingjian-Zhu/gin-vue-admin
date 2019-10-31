package datasource

import (
	"fmt"
	"log"

	"github.com/bingjian-zhu/gin-vue-admin/common/setting"
	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Db gormDB
type Db struct {
	Conn *gorm.DB
}

//Connect 初始化数据库配置
func (d *Db) Connect() error {
	var (
		dbType, dbName, user, pwd, host string
	)

	conf := setting.Config.Database
	dbType = conf.Type
	dbName = conf.Name
	user = conf.User
	pwd = conf.Password
	host = conf.Host

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, dbName))
	if err != nil {
		log.Fatal("connecting mysql error: ", err)
		return err
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	d.Conn = db

	log.Println("Connect Mysql Success")

	return nil
}

//DB 返回DB
func (d *Db) DB() *gorm.DB {
	return d.Conn
}

