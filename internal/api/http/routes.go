package http

import (
	"bbb-voting/internal/api/http/handler"
	"bbb-voting/internal/api/http/middleware"
	"bbb-voting/internal/core/ports"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, service ports.VoteService) {
	voteHandler := handler.NewVoteHandler(service)

	// API routes
	api := r.Group("/")
	{
		api.POST("/vote", middleware.CaptchaMiddleware(), voteHandler.CastVote)
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

	log.Println("Routes set up successfully")
}
