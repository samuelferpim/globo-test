package domain

import (
	"time"
)

type Vote struct {
	ID            string    `json:"id"`
	ParticipantID string    `json:"participant_id"`
	VoterID       string    `json:"voter_id"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	IPAddress     string    `json:"ip_address"`
	UserAgent     string    `json:"user_agent"`
	Region        string    `json:"region"`
	DeviceType    string    `json:"device_type"`
}

type VoteRedis struct {
	ID            string    `json:"id"`
	ParticipantID string    `json:"participant_id"`
	Timestamp     time.Time `json:"timestamp"`
}

type ParticipantResult struct {
	ParticipantID string `json:"participant_id"`
	VoteCount     int    `json:"vote_count"`
}

type DetailedResults struct {
	TotalVotes         int                 `json:"total_votes"`
	ParticipantResults []ParticipantResult `json:"participant_results"`
	VotesByHour        map[string]int      `json:"votes_by_hour"`
}
