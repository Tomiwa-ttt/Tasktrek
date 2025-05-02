package config

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() *mongo.Client {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatal("❌ Error loading .env file")
    }

    mongoURI := os.Getenv("MONGO_URI")
    clientOptions := options.Client().ApplyURI(mongoURI)

    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal("❌ MongoDB connection error:", err)
    }

    fmt.Println("✅ MongoDB connected successfully")
    DB = client
    return client
}




