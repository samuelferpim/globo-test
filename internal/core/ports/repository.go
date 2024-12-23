package ports

import (
	"bbb-voting/internal/core/domain"
	"context"
)

type VoteRepository interface {
	StoreInRedis(ctx context.Context, vote *domain.VoteRedis) error
	GetResultsByParticipant(ctx context.Context) (map[string]int, error)
	GetTotalVotes(ctx context.Context) (int, error)
	GetVotesByHour(ctx context.Context) (map[string]int, error)
}
