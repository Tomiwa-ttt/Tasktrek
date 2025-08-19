package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() *mongo.Client {
	// In production on Railway, .env won't exist. We rely on real env vars.
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("❌ MONGO_URI is empty (set it in Railway → Variables)")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("❌ MongoDB connection error:", err)
	}

	// Optionally ping:
	// if err := client.Ping(context.Background(), nil); err != nil {
	// 	log.Fatal("❌ MongoDB ping failed:", err)
	// }

	fmt.Println("✅ MongoDB connected successfully")
	DB = client
	return client
}
