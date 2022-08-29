package conn

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func openGorm(user, password, address, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		address,
		database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func openViperGorm() (*gorm.DB, error) {
	viper.SetDefault("mysql.user", "eshop")
	viper.SetDefault("mysql.password", "eshop!@#")
	viper.SetDefault("mysql.address", "localhost:3306")
	viper.SetDefault("mysql.database", "eshop")
	return openGorm(
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.address"),
		viper.GetString("mysql.database"),
	)
}

// ConnectToMySQL _
func ConnectToMySQL() *gorm.DB {
	db, err := openViperGorm()
	if err != nil {
		panic(err)
	}
	return db
}
