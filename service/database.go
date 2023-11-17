package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// 数据库处理

func (_ *DBService) InitDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://" + GetEvn("DB_HOST") + ":" + GetEvn("DB_PORT")).SetTimeout(10 * time.Second)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		utility.Log("DB Service ERROR:", err)
	}

	return client.Database(GetEvn("DB_DATA_BASE")) // 选择数据库
}

func (_ *DBService) ReadAllDB(server *mongo.Database, Collection string) []gin.H {
	// 读取数据库数据表内全部数据
	collection := server.Collection(Collection)
	cur, err := collection.Find(context.Background(), bson.D{}, options.Find())
	if err != nil {
		utility.Log("Read DB Data ERROR", err)
	}
	defer func(cur *mongo.Cursor, context context.Context) {
		err := cur.Close(context)
		if err != nil {
			utility.Log("Close BD ERROR:", err)
		}
	}(cur, context.Background())

	var data []gin.H

	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			utility.Log("Read All DB Data Decode ERROR", err)
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
		utility.Log("Read One DB Data Decode ERROR", err)
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
		utility.Log(err)
		return gin.H{}, false
	}
	return gin.H{"editCount": updateResult.ModifiedCount}, true
}
