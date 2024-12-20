package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bbb-voting/config"
	"bbb-voting/infra/database"
	"bbb-voting/infra/queue"
	httpdelivery "bbb-voting/internal/api/http"
	"bbb-voting/internal/core/service"
	"bbb-voting/internal/repository"
	"bbb-voting/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger.Init()
	defer logger.GetLogger().Sync()
	l := logger.GetLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		l.Fatal("Failed to load configuration", zap.Error(err))
	}

	redisDB, err := database.NewRedisDB(cfg.RedisURL)
	if err != nil {
		l.Fatal("Failed to connect to Redis", zap.Error(err))
	}
	defer redisDB.Close()

	rabbitMQ, err := queue.NewRabbitMQ(cfg.RabbitMQURL)
	if err != nil {
		l.Fatal("Failed to connect to RabbitMQ", zap.Error(err))
	}
	defer rabbitMQ.Close()

	voteRepo := repository.NewVoteRepository(redisDB)
	voteService := service.NewVoteService(voteRepo, rabbitMQ)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	httpdelivery.SetupRoutes(r, voteService)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	l.Info("Server started", zap.String("port", cfg.ServerPort))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	l.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		l.Fatal("Server forced to shutdown", zap.Error(err))
	}

	l.Info("Server exiting")
}
