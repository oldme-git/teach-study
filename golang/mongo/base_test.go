package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"runtime"
	"testing"
	"time"
)

func getConn() *mongo.Client {
	var (
		clientOptions = "mongodb://oldme:123456@192.168.10.43:27017/?authSource=oldme"
		ctx           = context.TODO()
	)
	ctx, cancelFun := context.WithTimeout(ctx, 3*time.Second)
	defer cancelFun()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(clientOptions))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}

func getCollection(client *mongo.Client, collection string) *mongo.Collection {
	return client.Database("oldme").Collection(collection)
}

func TestInsertOne(t *testing.T) {
	client := getConn()
	//user := bson.D{{"name", "wbbb"}, {"sex", 0}}
	user := bson.M{"name": "wbbb", "sex2": 0}
	res, err := getCollection(client, "users").InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	client.Disconnect(context.TODO())
	fmt.Println(runtime.NumGoroutine())
}

func TestFindOne(t *testing.T) {
	client := getConn()
	fitter := bson.D{{"sex", 0}}
	var res bson.D
	err := getCollection(client, "users").FindOne(context.TODO(), fitter).Decode(&res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestFind(t *testing.T) {
	client := getConn()
	fitter := bson.D{{"sex2", 0}}
	cursor, err := getCollection(client, "users").Find(context.TODO(), fitter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var res bson.M
		err := cursor.Decode(&res)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}

	client.Disconnect(context.TODO())

	// 检查游标错误
	if err := cursor.Err(); err != nil {
		panic(err)
	}
	fmt.Println(runtime.NumGoroutine())
}

func TestUpt(t *testing.T) {
	client := getConn()
	fitter := bson.D{{"sex2", 1}}
	upd := bson.D{{"$set", bson.D{{"name", "wb"}, {"sex", 3}, {"map", bson.M{"age": 20, "city": "nanjing"}}}}}
	one, err := getCollection(client, "users").UpdateOne(context.TODO(), fitter, upd)
	if err != nil {
		panic(err)
	}
	fmt.Println(one)
	client.Disconnect(context.TODO())
	fmt.Println(runtime.NumGoroutine())
}

func TestUpt2(t *testing.T) {
	client := getConn()
	fitter := bson.D{{"map.age", 22}}
	upd := bson.D{{"$set", bson.D{{"map", bson.M{"city": "beijing"}}}}}
	one, err := getCollection(client, "users").UpdateOne(context.TODO(), fitter, upd)
	if err != nil {
		panic(err)
	}
	fmt.Println(one)
	client.Disconnect(context.TODO())
	fmt.Println(runtime.NumGoroutine())
}

func TestDel(t *testing.T) {
	client := getConn()
	fitter := bson.D{{"map.city", "nanjing"}}
	res, err := getCollection(client, "users").DeleteOne(context.TODO(), fitter)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	client.Disconnect(context.TODO())
	fmt.Println(runtime.NumGoroutine())
}
