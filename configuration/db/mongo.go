package db

// import (
// 	"context"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func MongoConnect(host string, db string, username_mongo string, password_mongo string) *mongo.Database {

// 	// encodedUsername := url.QueryEscape(username_mongo)
// 	// encodedPassword := url.QueryEscape(password_mongo)
// 	// uri := fmt.Sprintf("mongodb://%s:%s@%s", encodedUsername, encodedPassword, host)
// 	uri := fmt.Sprintf("mongodb://%s", host)
// 	fmt.Println(uri)

// 	clientOptions := options.Client().ApplyURI(uri)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		fmt.Println("error mongo:", err)

// 	}

// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		fmt.Println("error mongo:", err)
// 		// log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")
// 	return client.Database(db)
// }
