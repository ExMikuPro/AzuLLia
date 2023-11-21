package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// 数据库处理

func InitDB() *mongo.Client {
	fmt.Println("mongodb://" + GetEvn("DB_HOST") + ":" + GetEvn("DB_PORT"))
	clientOptions := options.Client().ApplyURI("mongodb://" + GetEvn("DB_HOST") + ":" + GetEvn("DB_PORT")).SetTimeout(10 * time.Second)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		utilityFunction.Log("DB Service ERROR:", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Couldn't connect to the database:", err)
	}
	return client
}

func (s *DBService) ReadAllDB(Collection string, Data []gin.H) ([]gin.H, error) {
	// 读取数据库数据表内全部数据
	session, err := s.Client.StartSession()
	if err != nil {
		return nil, err
	}
	err = session.StartTransaction()
	if err != nil {
		return nil, err
	}
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		collection := s.Client.Database(GetEvn("DB_DATA_BASE")).Collection(Collection)
		cur, err := collection.Find(sessionContext, bson.D{}, options.Find())
		if err != nil {
			return err
		}
		for cur.Next(context.Background()) {
			var result bson.M
			err := cur.Decode(&result)
			if err != nil {
				return err
			}
			Data = append(Data, gin.H(result))
		}
		return nil
	})
	if err != nil {
		err := session.AbortTransaction(context.Background()) // 事务回滚
		if err != nil {
			return nil, err
		}
		return []gin.H{}, err
	}
	return Data, nil
}

func (s *DBService) ReadOneDB(Collection string, filter bson.D, Data gin.H) (gin.H, error) {
	session, err := s.Client.StartSession()
	if err != nil {
		return nil, err
	}
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		collection := s.Client.Database(GetEvn("DB_DATA_BASE")).Collection(Collection)
		err := collection.FindOne(sessionContext, filter).Decode(&Data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		err := session.AbortTransaction(context.Background())
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	return Data, nil
}

func (s *DBService) WriteOneDB(Collection string, Data any) error {
	session, err := s.Client.StartSession()
	if err != nil {
		return err
	}
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		collection := s.Client.Database(GetEvn("DB_DATA_BASE")).Collection(Collection)
		result, err := collection.InsertOne(sessionContext, Data)
		if err != nil {
			return err
		}
		_ = result.InsertedID.(primitive.ObjectID)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *DBService) UpdateOneDB(Collection string, filter bson.D, update bson.D) error {
	session, err := s.Client.StartSession()
	if err != nil {
		return err
	}
	err = session.StartTransaction()
	if err != nil {
		return err
	}
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		collection := s.Client.Database(GetEvn("DB_DATA_BASE")).Collection(Collection)
		_, err := collection.UpdateOne(sessionContext, filter, update)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *DBService) DeleteOneDB(Collection string, filter bson.D) error {
	session, err := s.Client.StartSession()
	if err != nil {
		return err
	}
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		collection := s.Client.Database(GetEvn("DB_DATA_BASE")).Collection(Collection)
		_, err := collection.DeleteOne(sessionContext, filter)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
