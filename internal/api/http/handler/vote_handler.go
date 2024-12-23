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

// CastVote godoc
// @Summary Cast a vote
// @Description Cast a vote for a BBB participant
// @Tags votes
// @Accept json
// @Produce json
// @Param vote body domain.Vote true "Vote details"
// @Success 200 {object} map[string]interface{} "Successfully cast vote"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Failed to cast vote"
// @Router /vote [post]
func (h *VoteHandler) CastVote(c *gin.Context) {
	var vote domain.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.CastVote(c.Request.Context(), &vote); err != nil {
		log.Printf("Failed to cast vote: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast vote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote cast successfully"})
}

// GetTotalVotes godoc
// @Summary Get total votes
// @Description Get the total number of votes cast
// @Tags votes
// @Produce json
// @Success 200 {object} map[string]interface{} "Total votes"
// @Failure 500 {object} map[string]interface{} "Failed to get total votes"
// @Router /votes/total [get]
func (h *VoteHandler) GetTotalVotes(c *gin.Context) {
	total, err := h.service.GetTotalVotes(c.Request.Context())
	if err != nil {
		log.Printf("Failed to get total votes: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total votes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_votes": total})
}

// GetDetailedResults godoc
// @Summary Get detailed voting results
// @Description Get detailed results of the voting, including votes per participant
// @Tags votes
// @Produce json
// @Success 200 {object} map[string]interface{} "Detailed voting results"
// @Failure 500 {object} map[string]interface{} "Failed to get detailed results"
// @Router /votes/detailed [get]
func (h *VoteHandler) GetDetailedResults(c *gin.Context) {
	votes, err := h.service.GetDetailedResults(c.Request.Context())
	if err != nil {
		log.Printf("Failed to get detailed results: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get detailed results"})
		return
	}
	c.JSON(http.StatusOK, votes)
}

// GetVotesByHour godoc
// @Summary Get votes by hour
// @Description Get the number of votes cast per hour
// @Tags votes
// @Produce json
// @Success 200 {object} map[string]interface{} "Votes per hour"
// @Failure 500 {object} map[string]interface{} "Failed to get votes by hour"
// @Router /votes/by-hour [get]
func (h *VoteHandler) GetVotesByHour(c *gin.Context) {
	votesPerHour, err := h.service.GetVotesByHour(c.Request.Context())
	if err != nil {
		log.Printf("Failed to get votes by hour: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get votes by hour"})
		return
	}
	c.JSON(http.StatusOK, votesPerHour)
}
