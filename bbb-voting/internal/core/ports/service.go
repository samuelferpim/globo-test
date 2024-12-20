package ports

import (
	"bbb-voting/internal/core/domain"
	"context"
)

type VoteService interface {
	CastVote(ctx context.Context, vote *domain.Vote) error
	GetDetailedResults(ctx context.Context) (domain.DetailedResults, error)
	GetTotalVotes(ctx context.Context) (int, error)
	GetVotesByHour(ctx context.Context) (map[string]int, error)
}
