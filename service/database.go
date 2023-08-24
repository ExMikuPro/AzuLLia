package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 数据库处理

func InitDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI(DBAddress)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("DB Service ERROR:", err)
	}
	return client.Database(DBDataBase) // 选择数据库
}

func ReadAllDB(server *mongo.Database, Collection string) []gin.H {
	// todo 做性能优化
	// 读取数据库数据表内全部数据
	collection := server.Collection(Collection)

	filter := bson.D{}
	option := options.Find()

	cur, err := collection.Find(context.Background(), filter, option)
	if err != nil {
		fmt.Println("Read DB Data ERROR", err)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			fmt.Println("Close BD ERROR:", err)
		}
	}(cur, context.Background())

	var data []gin.H

	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println("Read All DB Data Decode ERROR", err)
		}
		data = append(data, gin.H(result))
	}
	return data
}

func ReadOneDB(server *mongo.Database, Collection string, filter bson.D) gin.H { // todo 性能优化
	collection := server.Collection(Collection)
	result := collection.FindOne(context.Background(), filter)
	var data = gin.H{}
	err := result.Decode(&data)
	if err != nil {
		fmt.Println("Read One DB Data Decode ERROR", err)
	}
	return data
}
