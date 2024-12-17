package config

import (
    "fmt"
    "log"
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// ConnectToDatabase 初始化数据库连接
func ConnectToDatabase() {
    var err error
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
        os.Getenv("DB_USER"), 
        os.Getenv("DB_PASSWORD"), 
        os.Getenv("DB_HOST"), 
        os.Getenv("DB_PORT"), 
        os.Getenv("DB_NAME"))
    
    DB, err = gorm.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    log.Println("Successfully connected to the database.")
}