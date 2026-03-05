// db/mongo.go
package db

import (
    "context"
    "fmt"
    "os"
    "strconv"
    "time"

    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.uber.org/zap"
)

type MongoConfig struct {
    URI         string
    DBName      string
    MaxPoolSize uint64
    MinPoolSize uint64
}

type MongoConnection struct {
    Client   *mongo.Client
    Database *mongo.Database
}

func NewMongoConfig() MongoConfig {
    return MongoConfig{
        URI:         os.Getenv("MONGO_URI"),
        DBName:      os.Getenv("MONGO_DB_NAME"),
        MaxPoolSize: parseUint(os.Getenv("MONGO_MAX_POOL_SIZE"), 25),
        MinPoolSize: parseUint(os.Getenv("MONGO_MIN_POOL_SIZE"), 5),
    }
}

func NewMongoConnection(cfg MongoConfig, sugar *zap.SugaredLogger) (*MongoConnection, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().
        ApplyURI(cfg.URI).
        SetServerAPIOptions(serverAPI).
        SetMaxPoolSize(cfg.MaxPoolSize).
        SetMinPoolSize(cfg.MinPoolSize).
        SetMaxConnIdleTime(5 * time.Minute)

    client, err := mongo.Connect(opts)
    if err != nil {
        return nil, fmt.Errorf("failed to create MongoDB client: %w", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
    }

    sugar.Infow("Connected to MongoDB",
        "dbName",      cfg.DBName,
        "maxPoolSize", cfg.MaxPoolSize,
        "minPoolSize", cfg.MinPoolSize,
    )

    return &MongoConnection{
        Client:   client,
        Database: client.Database(cfg.DBName),
    }, nil
}

func parseUint(val string, defaultVal uint64) uint64 {
    if val == "" {
        return defaultVal
    }
    parsed, err := strconv.ParseUint(val, 10, 64)
    if err != nil {
        return defaultVal
    }
    return parsed
}