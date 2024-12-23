package http

import (
	"bbb-voting/internal/api/http/handler"
	middleware "bbb-voting/internal/api/http/middleware"
	"bbb-voting/internal/core/ports"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRoutes(r *gin.Engine, service ports.VoteService, logger *zap.Logger) {
	voteHandler := handler.NewVoteHandler(service, logger)

	r.Use(func(c *gin.Context) {
		start := time.Now()

		c.Next()

		logger.Info("Request",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", time.Since(start)),
		)
	})

	// API routes
	api := r.Group("/")
	{
		api.POST("/vote", middleware.CaptchaMiddleware(logger), voteHandler.CastVote)
		api.GET("/result", voteHandler.GetDetailedResults)
		api.GET("/total-votes", voteHandler.GetTotalVotes)
		api.GET("/votes-by-hour", voteHandler.GetVotesByHour)
	}

	// Static files
	r.Static("/static", "./web/static")

	// HTML route
	r.GET("/", func(c *gin.Context) {
		c.File("./web/templates/index.html")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	})

	logger.Info("Routes set up successfully")

}
