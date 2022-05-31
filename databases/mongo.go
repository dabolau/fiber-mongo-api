package databases

import (
	"context"
	"fiber-mongo-api/configs"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 数据库连接字符串
var mongoURI = configs.EnvMongoURI()

// 数据库名称
var mongoDbName = "fiber-mongo-api"

// 数据库连接
// https://www.mongodb.com/blog/post/quick-start-golang--mongodb--starting-and-setup
func Connect() *mongo.Client {
	// 创建客户端并绑定资源地址
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	// 创建上下文并可以取消控制
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 通过客户端连接数据库
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)

	}
	// 测试数据库连接是否成功
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)

	}
	return client
}

// 数据库客户端
var MongoClient = Connect()

// 获取数据库集合
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(mongoDbName).Collection(collectionName)
}
