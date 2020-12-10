package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Member struct {
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/bitkingreturns"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("testing").Collection("numbers")
	start := time.Now()
	cur, err := collection.Find(
		context.Background(),
		bson.M{
			// "roi": bson.M{
			// 	"$lte": 10.07 - 0.30625,
			// },
			"roundId": "274f7ac0",
		},
	)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed.Milliseconds())
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.Background()) {
		//Create a value into which the single document can be decoded
		result := struct {
			Code    string
			roundId string
		}{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//Close the cursor once finished
	cur.Close(context.TODO())
}
