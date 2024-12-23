package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bbb-voting/config"

	httpdelivery "bbb-voting/internal/api/http"
	"bbb-voting/internal/core/ports"
	serviceRepo "bbb-voting/internal/core/repository"
	"bbb-voting/internal/core/service"
	"bbb-voting/internal/infra/queue"
	"bbb-voting/internal/infra/redis"
	"bbb-voting/internal/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func main() {
	c := buildContainer()

	err := c.Invoke(func(srv *http.Server, l *zap.Logger) {
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				l.Fatal("Failed to start server", zap.Error(err))
			}
		}()

		l.Info("Server started", zap.String("addr", srv.Addr))

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
	})

	if err != nil {
		log.Fatal(err)
	}
}

func buildContainer() *dig.Container {
	c := dig.New()

	// Provide logger
	c.Provide(func() (*zap.Logger, error) {
		return zap.NewProduction()
	})

	// Provide configuration
	c.Provide(config.LoadConfig)

	// Provide Redis client
	c.Provide(func(cfg *config.Config, logger *zap.Logger) (repository.RedisClient, error) {
		return redis.NewRedisClient(cfg.RedisURL, logger)
	})

	// Provide AMQP connection
	c.Provide(func(cfg *config.Config) (ports.AMQPConnection, error) {
		return queue.NewRabbitMQConnection(cfg.RabbitMQURL)
	})

	// Provide RabbitMQ service
	c.Provide(func(conn ports.AMQPConnection) ports.QueueService {
		return queue.NewRabbitMQ(conn)
	})

	// Provide VoteRepository
	c.Provide(func(redisClient repository.RedisClient) ports.VoteRepository {
		return serviceRepo.NewVoteRepository(redisClient)
	})

	// Provide VoteService
	c.Provide(func(repo ports.VoteRepository, queue ports.QueueService) ports.VoteService {
		return service.NewVoteService(service.VoteServiceOptions{
			Repo:  repo,
			Queue: queue,
		})
	})

	// Provide Gin engine
	c.Provide(func() *gin.Engine {
		gin.SetMode(gin.ReleaseMode)
		r := gin.New()
		r.Use(gin.Recovery(), gin.Logger())
		return r
	})

	// Provide HTTP server
	c.Provide(func(cfg *config.Config, r *gin.Engine) *http.Server {
		return &http.Server{
			Addr:    ":" + cfg.ServerPort,
			Handler: r,
		}
	})

	// Setup routes
	c.Invoke(func(r *gin.Engine, vs ports.VoteService) {
		httpdelivery.SetupRoutes(r, vs)
	})

	return c
}
