package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 获取环境变量的数据库连接字符串
// mongodb://username:password@localhost:27017
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return fmt.Sprintf("mongodb://%s:%s@%s", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_ADDRESS"))
}
