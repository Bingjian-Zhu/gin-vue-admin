package datasource

import "github.com/jinzhu/gorm"

//IDb 数据库接口定义
type IDb interface {
	//Connect 初始化数据库配置
	Connect() error
	//DB 返回DB
	DB() *gorm.DB
}
