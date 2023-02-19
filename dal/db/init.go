package db

import (
	"github.com/simple/douyin/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if err = db.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&UserRaw{}, &VideoRaw{}, &FavoriteRaw{}, &CommentRaw{}, &RelationRaw{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(constants.MySQLMaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(constants.MySQLMaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(constants.MySQLConnMaxLifetime)

}
