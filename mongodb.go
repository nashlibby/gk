package gk

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func NewMongodb(dsn string) (*Mongodb, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	return &Mongodb{Client: client}, nil
}

func (m *Mongodb) Collection(database, collection string) *Mongodb {
	m.Coll = m.Client.Database(database).Collection(collection)
	return m
}

// 创建索引
func (m *Mongodb) CreateIndex(index mongo.IndexModel, opts ...*options.CreateIndexesOptions) (res string, err error) {
	return m.Coll.Indexes().CreateOne(context.TODO(), index, opts...)
}

// 查找单条记录
func (m *Mongodb) Find(filter bson.D) (res string, err error) {
	var result bson.M
	err = m.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		return
	}

	res = string(jsonData)

	return
}

// 获取列表
func (m *Mongodb) Get(query bson.D) (res string, err error) {
	cursor, err := m.Coll.Find(context.TODO(), query)
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(cursor)
	if err != nil {
		return
	}
	res = string(jsonData)

	return
}

// 插入一条数据
func (m *Mongodb) Insert(doc bson.D) (res string, err error) {
	result, err := m.Coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		return
	}
	res = string(jsonData)

	return
}

// 插入多条数据
func (m *Mongodb) InsertMany(docs []interface{}) (res string, err error) {
	results, err := m.Coll.InsertMany(context.TODO(), docs)
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(results)
	if err != nil {
		return
	}
	res = string(jsonData)

	return
}

// 按id更新一条数据
func (m *Mongodb) UpdateOneById(id string, doc bson.D) (res string, err error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	results, err := m.Coll.UpdateByID(context.TODO(), objectId, doc)
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(results)
	if err != nil {
		return
	}
	res = string(jsonData)

	return
}

// 按条件更新多条数据
func (m *Mongodb) UpdateManyByFilter(filter, doc bson.D) (res string, err error) {
	results, err := m.Coll.UpdateMany(context.TODO(), filter, doc)
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(results)
	if err != nil {
		return
	}
	res = string(jsonData)

	return
}

// 按条件替换一条记录
func (m *Mongodb) ReplaceOneByFilter(filter, doc bson.D) (res string, err error) {
	results, err := m.Coll.ReplaceOne(context.TODO(), filter, doc)
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(results)
	if err != nil {
		return
	}
	res = string(jsonData)

	return
}

// 按条件删除多条记录
func (m *Mongodb) DeleteManyByFilter(filter bson.D) (res string, err error) {
	results, err := m.Coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(results)
	if err != nil {
		return
	}
	res = string(jsonData)

	return
}

// 断开连接
func (m *Mongodb) DisConnect() {
	if err := m.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
