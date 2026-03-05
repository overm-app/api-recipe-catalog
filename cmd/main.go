package main

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/overm-app/api-recipe-catalog/internal/infrastructure/db"
)

func main() {
	// Setup logger
	sugar := setupLogger()
	defer sugar.Sync()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		sugar.Warnw("No .env file found, using environment variables or defaults")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081"
		sugar.Warnw("SERVER_PORT not set, defaulting to 8081")
	}

	// jwtExpiration := 24 * time.Hour
	// if exp := os.Getenv("JWT_EXPIRATION_HOURS"); exp != "" {
	// 	if parsed, err := time.ParseDuration(exp + "h"); err == nil {
	// 		jwtExpiration = parsed
	// 	} else {
	// 		sugar.Warnw("Invalid JWT_EXPIRATION_HOURS, defaulting to 24h", "value", exp)
	// 	}
	// }

	// Set server timezone
	setupTimezone(sugar)

	// Connect to databasepackage cmd
	mongoCfg := db.MongoConfig{
		URI:    os.Getenv("MONGO_URI"),
		DBName: os.Getenv("MONGO_DB_NAME"),
		MaxPoolSize: parseUint(os.Getenv("MONGO_MAX_POOL_SIZE"), 25),
		MinPoolSize: parseUint(os.Getenv("MONGO_MIN_POOL_SIZE"), 5),
	}

	mongoConn, err := db.NewMongoConnection(mongoCfg, sugar)
	if err != nil {
		sugar.Errorw("Failed to connect to MongoDB", "error", err)
		os.Exit(1)
	}
	defer mongoConn.Client.Disconnect(context.Background())

	// Initialize repositories

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		sugar.Errorw("JWT_SECRET is not set")
		os.Exit(1)
	}

	// Initialize services

	// Initialize use cases

	// Initialize API handlers

	// Setup and start server
	// r := api.NewRouter(authHandler, userHandler, jwtService, sugar)
	// engine := r.SetupRouter(sugar)

	sugar.Infow("Starting server", "port", port)

	// if err := engine.Run(":" + port); err != nil {
	// 	sugar.Errorw("Failed to start server",
	// 		"error", err,
	// 		"port", port,
	// 	)
	// 	os.Exit(1)
	// }
}

func setupLogger() *zap.SugaredLogger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	var level zapcore.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		level = zapcore.InfoLevel
	}

	logger, _ := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    encoderCfg,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()

	logger = logger.WithOptions(zap.WithCaller(true), zap.AddStacktrace(zapcore.FatalLevel))
	sugar := logger.Sugar()
	return sugar
}

func setupTimezone(sugar *zap.SugaredLogger) {
	timezone := os.Getenv("SERVER_TIMEZONE")
	if timezone == "" {
		timezone = "UTC"
	}
	loc, err := time.LoadLocation(timezone)
	sugar.Infow("Setting server timezone", "timezone", timezone)
	if err != nil {
		sugar.Errorw("Failed to set timezone", "timezone", timezone, "error", err)
		loc = time.UTC
	}
	time.Local = loc
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