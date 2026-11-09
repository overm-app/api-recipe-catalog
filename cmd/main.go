package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/overm-app/api-recipe-catalog/internal/infrastructure/db"
	"github.com/overm-app/api-recipe-catalog/internal/infrastructure/repository"
	"github.com/overm-app/api-recipe-catalog/internal/infrastructure/service"
	"github.com/overm-app/api-recipe-catalog/internal/interface/api"
	"github.com/overm-app/api-recipe-catalog/internal/interface/api/handlers"
	"github.com/overm-app/api-recipe-catalog/internal/usecase"
)

func main() {
	// Setup logger
	sugar := setupLogger()
	defer sugar.Sync()

	currentEnv := os.Getenv("ENV")
	if currentEnv != "develop" && currentEnv != "QA" && currentEnv != "prod" {
		if err := godotenv.Load(); err != nil {
			sugar.Errorw("No .env file found, relying on system environment variables")
		}
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8082"
		sugar.Warnw("SERVER_PORT not set, defaulting to 8082")
	}

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
	recipeRepo := repository.NewRecipeRepository(mongoConn)

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		sugar.Errorw("JWT_SECRET is not set")
		os.Exit(1)
	}

	// Initialize services
	jwtService := service.NewJWTService(jwtSecret)

	// Initialize use cases
	recipeUseCase := usecase.NewRecipeUseCase(recipeRepo)

	// Initialize API handlers
	recipeHandler := handlers.NewRecipeHandler(recipeUseCase, jwtService, sugar)
	healthHandler := handlers.NewHealthHandler(mongoConn, sugar)

	// Setup and start server
	r := api.NewRouter(recipeHandler, healthHandler, jwtService, sugar)
	engine := r.SetupRouter(sugar)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: engine,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	serverErr := make(chan error, 1)
	go func() {
		sugar.Infow("Starting server", "port", port)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		sugar.Errorw("Server failed", "error", err)
		os.Exit(1)
	case <-ctx.Done():
	}

	stop()

	sugar.Infow("Shutdown signal received")
	healthHandler.SetShuttingDown()

	time.Sleep(5 * time.Second)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		sugar.Errorw("Graceful shutdown failed", "error", err)
	}

	sugar.Infow("Server stopped")
}

func setupLogger() *zap.SugaredLogger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		zapcore.ISO8601TimeEncoder(t.UTC(), enc)
	}
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