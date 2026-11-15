package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/overm-app/api-recipe-catalog/internal/config"
	"github.com/overm-app/api-recipe-catalog/internal/infrastructure/db"
	"github.com/overm-app/api-recipe-catalog/internal/infrastructure/repository"
	"github.com/overm-app/api-recipe-catalog/internal/infrastructure/service"
	"github.com/overm-app/api-recipe-catalog/internal/interface/api"
	"github.com/overm-app/api-recipe-catalog/internal/interface/api/handlers"
	"github.com/overm-app/api-recipe-catalog/internal/usecase"
)

func main() {
	// Setup logger
	sugar, level := newLogger()
	defer sugar.Sync()

	path := os.Getenv("CONFIG_FILE")
	if path == "" {
		path = "/etc/config/config.yaml"
	}
	cfg, err := config.Load(path)
	if err != nil {
		sugar.Errorw("invalid configuration", "error", err)
		os.Exit(1)
	}

	var lvl zapcore.Level
	if err := lvl.UnmarshalText([]byte(cfg.Log.Level)); err != nil {
		sugar.Errorw("invalid log.level", "value", cfg.Log.Level, "error", err)
		os.Exit(1)
	}
	level.SetLevel(lvl)
	sugar = sugar.With("env", cfg.Env)

	port := cfg.Server.Port

	mongoCfg := db.MongoConfig{
		URI:    cfg.MongoURI,
		DBName: cfg.Database.DBName,
		MaxPoolSize: cfg.Database.Pool.Max,
		MinPoolSize: cfg.Database.Pool.Min,
	}

	mongoConn, err := db.NewMongoConnection(mongoCfg, sugar)
	if err != nil {
		sugar.Errorw("Failed to connect to MongoDB", "error", err)
		os.Exit(1)
	}
	defer mongoConn.Client.Disconnect(context.Background())

	// Initialize repositories
	recipeRepo := repository.NewRecipeRepository(mongoConn)

	jwtSecret := []byte(cfg.JWTSecret)
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

func newLogger() (*zap.SugaredLogger, zap.AtomicLevel) {
	level := zap.NewAtomicLevelAt(zapcore.InfoLevel)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		zapcore.ISO8601TimeEncoder(t.UTC(), enc)
	}
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	logger, err := zap.Config{
		Level:            level,
		Encoding:         "json",
		EncoderConfig:    encoderCfg,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(zap.WithCaller(true), zap.AddStacktrace(zapcore.FatalLevel))
	if err != nil {
		panic(err)
	}
	return logger.With(zap.String("service", "api-menu-shuffle")).Sugar(), level
}