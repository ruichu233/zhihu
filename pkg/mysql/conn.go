package mysql

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"

	"time"
)

type Conf struct {
	UserName string `json:"user_name,omitempty" yaml:"user_name"`
	Password string `json:"password,omitempty" yaml:"password"`
	Host     string `json:"host,omitempty" yaml:"host"`
	Port     string `json:"port,omitempty" yaml:"port"`
	Database string `json:"database,omitempty" yaml:"database"`
	Charset  string `json:"charset,omitempty" yaml:"charset"`
}

func InitMysql(opt *Conf) *gorm.DB {
	//配置MySQL连接参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", opt.UserName, opt.Password, opt.Host, opt.Port, opt.Database, opt.Charset)
	db, err := database(dsn)
	if err != nil {
		log.Fatalf("mysql connect error: %v", err)
	}
	return db
}

func database(connString string) (*gorm.DB, error) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connString, // DSN data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	return db, err
}
