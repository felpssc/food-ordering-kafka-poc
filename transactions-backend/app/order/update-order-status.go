package order

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func UpdateOrderStatus(status string, document_id string) {
	// Set client options
	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI(os.Getenv("MONGO_DSN")))

	if err != nil {
		fmt.Println(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		fmt.Println(err)
	}

	updateQuery := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: status},
		}},
	}

	ordersCollection := client.Database("nest").Collection("orders")

	objID, _ := primitive.ObjectIDFromHex(document_id)

	r, err := ordersCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		updateQuery,
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Updated", r.ModifiedCount, "order(s) status on database. ID: ", document_id)
}
