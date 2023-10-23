package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 数据库处理

func (_ *DBService) InitDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI(DBAddress)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("DB Service ERROR:", err)
	}
	return client.Database(DBDataBase) // 选择数据库
}

func (_ *DBService) ReadAllDB(server *mongo.Database, Collection string) []gin.H {
	// todo 做性能优化
	// 读取数据库数据表内全部数据
	collection := server.Collection(Collection)

	filter := bson.D{}
	option := options.Find()

	cur, err := collection.Find(context.Background(), filter, option)
	if err != nil {
		fmt.Println("Read DB Data ERROR", err)
	}
	defer func(cur *mongo.Cursor, context context.Context) {
		err := cur.Close(context)
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

func (_ *DBService) ReadOneDB(server *mongo.Database, Collection string, filter bson.D) (gin.H, bool) { // todo 性能优化
	collection := server.Collection(Collection)
	result := collection.FindOne(context.Background(), filter)
	var data = gin.H{}
	err := result.Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments { // 判断如果没有文档
			return gin.H{}, false
		}
		fmt.Println("Read One DB Data Decode ERROR", err)
		return gin.H{}, false
	}
	return data, true
}

func (_ *DBService) WriteOneDB(server *mongo.Database, Collection string, Data any) (string, bool) {
	collection := server.Collection(Collection)
	one, err := collection.InsertOne(context.TODO(), Data)
	insertedID := one.InsertedID.(primitive.ObjectID)
	if err != nil {
		return insertedID.Hex(), false
	}
	return insertedID.Hex(), true
}

func (_ *DBService) UpdateOneDB(server *mongo.Database, Collection string, filter bson.D, update bson.D) (gin.H, bool) {
	collection := server.Collection(Collection)
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
		return gin.H{}, false
	}
	return gin.H{"editCount": updateResult.ModifiedCount}, true
}
