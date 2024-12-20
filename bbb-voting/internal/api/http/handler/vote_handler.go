package handler

import (
	"log"
	"net/http"

	"bbb-voting/internal/core/domain"
	"bbb-voting/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type VoteHandler struct {
	service ports.VoteService
}

func NewVoteHandler(service ports.VoteService) *VoteHandler {
	return &VoteHandler{service: service}
}

func (h *VoteHandler) CastVote(c *gin.Context) {
	var vote domain.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CastVote(c.Request.Context(), &vote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast vote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote cast successfully"})
}

func (h *VoteHandler) GetTotalVotes(c *gin.Context) {
	total, err := h.service.GetTotalVotes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total votes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_votes": total})
}

func (h *VoteHandler) GetDetailedResults(c *gin.Context) {
	votes, err := h.service.GetDetailedResults(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get votes by participant"})
		return
	}
	c.JSON(http.StatusOK, votes)
}

func (h *VoteHandler) GetVotesByHour(c *gin.Context) {
	votesPerHour, err := h.service.GetVotesByHour(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get votes by hour"})
		return
	}
	c.JSON(http.StatusOK, votesPerHour)
}
