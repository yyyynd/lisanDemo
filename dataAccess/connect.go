package dataAccess

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// The following default values are reserved for the test environment
var (
	DB_NAME = "lisandb"
	USER    = "menghy0523"
	PASSWD  = "menghy0523"
)

func InitConnection(user string, passwd string, address string, dbname string) (*gorm.DB, error) {
	db, err := initConnection(user, passwd, address, dbname)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initConnection(user string, passwd string, address string, dbname string) (*gorm.DB, error) {
	dsn := passwd + ":" + user + "@tcp(127.0.0.1:3306)/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateConnection(dbname string) (*gorm.DB, error) {
	dsn := "menghy0523:menghy0523@tcp(127.0.0.1:3306)/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
