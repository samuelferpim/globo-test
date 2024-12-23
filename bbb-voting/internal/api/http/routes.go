package http

import (
	"bbb-voting/internal/api/http/handler"
	middleware "bbb-voting/internal/api/http/middleware"
	"bbb-voting/internal/core/ports"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, service ports.VoteService) {
	voteHandler := handler.NewVoteHandler(service)

	r.POST("/vote", middleware.CaptchaMiddleware(), voteHandler.CastVote)
	r.GET("/result", voteHandler.GetDetailedResults)
	r.GET("/total-votes", voteHandler.GetTotalVotes)
	r.GET("/votes-by-hour", voteHandler.GetVotesByHour)

	r.Static("/static", "../web/static")

	r.GET("/", func(c *gin.Context) {
		c.File("../web/templates/index.html")
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
		})
	})
}
