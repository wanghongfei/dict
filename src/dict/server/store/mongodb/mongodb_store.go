package mongodb

import (
	"dict/client/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongodbStore struct {
	connection *mgo.Session
}

// 新建实例
// host: redis地址
func NewMongodbStore(host string) (*MongodbStore, error) {
	// 连接数据库
	session, err := mgo.Dial(host)
	if nil != err {
		return nil, err
	}

	instance := &MongodbStore{
		connection: session,
	}

	return instance, nil
}

// 关闭连接
func (me *MongodbStore) Close() {
	me.connection.Close()
}

// 将单词写入mongodb
func (me *MongodbStore) Save(word *model.Word) error {
	conn := me.connection.DB("test").C("vocabulary")

	err := conn.Insert(word)
	if nil != err {
		return err
	}

	return nil
}

// 查询单词
func (me *MongodbStore) Load(word string) *model.Word {
	conn := me.connection.DB("test").C("vocabulary")

	result := &model.Word{}
	conn.Find(bson.M{
		"literal": word,
	}).One(result)

	return result
}

