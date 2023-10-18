package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/xishengcai/cloud/pkg/setting"
)

func InitMysql() {
	m := setting.Config.Mysql
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.DBName,
		m.Charset,
		m.ParseTime))
	if err != nil {
		panic(err)
	}
	conn.SetMaxIdleConns(m.MaxIdleConnections)
	conn.SetMaxOpenConns(m.MaxOpenConnections)
	setting.DB, err = gorm.Open(
		mysql.New(mysql.Config{Conn: conn}),
		&gorm.Config{
			Logger:                 newLogger,
			NamingStrategy:         schema.NamingStrategy{SingularTable: true},
			SkipDefaultTransaction: true,
		})

	if err != nil {
		panic(err)
	}

	setting.DB.Debug()
}

var (
	newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
)
