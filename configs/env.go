package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 获取环境变量中的数据库连接字符串
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGOURI")
}

// 获取环境变量中的数据库名称
func EnvMongoDbName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("DBNAME")
}
