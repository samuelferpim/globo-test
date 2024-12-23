package handler

import (
	"net/http"

	"bbb-voting/internal/core/domain"
	"bbb-voting/internal/core/ports"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VoteHandler struct {
	service ports.VoteService
	logger  *zap.Logger
}

func NewVoteHandler(service ports.VoteService, logger *zap.Logger) *VoteHandler {
	return &VoteHandler{service: service, logger: logger}
}

func (h *VoteHandler) CastVote(c *gin.Context) {
	var vote domain.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		h.logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.CastVote(c.Request.Context(), &vote); err != nil {
		h.logger.Error("Failed to cast vote", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast vote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote cast successfully"})
}

func (h *VoteHandler) GetTotalVotes(c *gin.Context) {
	total, err := h.service.GetTotalVotes(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get total votes", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total votes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_votes": total})
}

func (h *VoteHandler) GetDetailedResults(c *gin.Context) {
	votes, err := h.service.GetDetailedResults(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get detailed results", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get detailed results"})
		return
	}
	c.JSON(http.StatusOK, votes)
}

func (h *VoteHandler) GetVotesByHour(c *gin.Context) {
	votesPerHour, err := h.service.GetVotesByHour(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get votes by hour", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get votes by hour"})
		return
	}
	c.JSON(http.StatusOK, votesPerHour)
}
