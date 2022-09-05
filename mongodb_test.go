package gk

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestNewMongodb(t *testing.T) {
	client, err := NewMongodb("mgdb://root:example@127.0.0.1:27017")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(client)
	client.DisConnect()
}

// 测试查找单条记录
func TestMongodb_Find(t *testing.T) {
	client, err := NewMongodb("mgdb://root:example@127.0.0.1:27017")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := client.Collection("supplier", "persons").Find(bson.D{
		{"name", "leo"},
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(res)

	client.DisConnect()
}

// 测试单条记录插入
func TestMongodb_Insert(t *testing.T) {
	client, err := NewMongodb("mgdb://root:example@127.0.0.1:27017")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := client.Collection("supplier", "persons").Insert(bson.D{
		{"name", "nash"},
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(res)

	client.DisConnect()
}

// 测试多条记录插入
func TestMongodb_InsertMany(t *testing.T) {
	client, err := NewMongodb("mgdb://root:example@127.0.0.1:27017")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := client.Collection("supplier", "persons").InsertMany([]interface{}{
		bson.D{
			{"name", "nash"},
		},
		bson.D{
			{"name", "leo"},
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(res)

	client.DisConnect()
}

// 测试根据id更新单条记录
func TestMongodb_UpdateOneById(t *testing.T) {
	client, err := NewMongodb("mgdb://root:example@127.0.0.1:27017")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := client.Collection("supplier", "persons").UpdateOneById("62bbc3648dcd43e69e4b84e0", bson.D{
		{"$set", bson.D{
			{"name", "libby"},
		}},
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(res)

	client.DisConnect()
}

// 测试根据条件更新多条记录
func TestMongodb_UpdateManyByFilter(t *testing.T) {
	client, err := NewMongodb("mgdb://root:example@127.0.0.1:27017")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := client.Collection("supplier", "persons").UpdateManyByFilter(bson.D{
		{"name", "nash"},
	}, bson.D{
		{"$set", bson.D{
			{"age", 18},
		}},
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(res)

	client.DisConnect()
}
