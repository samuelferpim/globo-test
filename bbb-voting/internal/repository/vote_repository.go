package repository

import (
	"context"
	"fmt"
	"strings"

	"bbb-voting/internal/core/domain"
	"bbb-voting/internal/core/ports"
)

type voteRepository struct {
	redis ports.RedisRepository
}

func NewVoteRepository(redis ports.RedisRepository) ports.VoteRepository {
	return &voteRepository{
		redis: redis,
	}
}

func (r *voteRepository) StoreInRedis(ctx context.Context, vote *domain.VoteRedis) error {
	pipe := r.redis.Pipeline()

	// Store individual vote
	voteKey := fmt.Sprintf("vote:%s", vote.ID)
	pipe.HMSet(ctx, voteKey, map[string]interface{}{
		"id":             vote.ID,
		"participant_id": vote.ParticipantID,
		"timestamp":      vote.Timestamp.Unix(),
	})

	// Increment vote count for participant
	participantCountKey := fmt.Sprintf("count:participant:%s", vote.ParticipantID)
	pipe.Incr(ctx, participantCountKey)

	// Increment total vote count
	pipe.Incr(ctx, "count:total")

	// Increment vote count for the current hour
	hourKey := fmt.Sprintf("count:hour:%s", vote.Timestamp.Format("2006010215"))
	pipe.Incr(ctx, hourKey)

	return pipe.Exec(ctx)
}

func (r *voteRepository) GetResultsByParticipant(ctx context.Context) (map[string]int, error) {
	var cursor uint64
	results := make(map[string]int)

	for {
		scanCmd := r.redis.Scan(ctx, cursor, "count:participant:*", 10)
		keys, cursor, err := scanCmd.Result()
		if err != nil {
			return nil, err
		}

		for _, key := range keys {
			participantID := strings.TrimPrefix(key, "count:participant:")
			countCmd := r.redis.Get(ctx, key)
			count, err := countCmd.Int()
			if err != nil {
				return nil, err
			}
			results[participantID] = count
		}

		if cursor == 0 {
			break
		}
	}

	return results, nil
}

func (r *voteRepository) GetTotalVotes(ctx context.Context) (int, error) {
	cmd := r.redis.Get(ctx, "count:total")
	return cmd.Int()
}

func (r *voteRepository) GetVotesByHour(ctx context.Context) (map[string]int, error) {
	var cursor uint64
	results := make(map[string]int)

	for {
		scanCmd := r.redis.Scan(ctx, cursor, "count:hour:*", 10)
		keys, cursor, err := scanCmd.Result()
		if err != nil {
			return nil, err
		}

		for _, key := range keys {
			hour := strings.TrimPrefix(key, "count:hour:")
			countCmd := r.redis.Get(ctx, key)
			count, err := countCmd.Int()
			if err != nil {
				return nil, err
			}
			results[hour] = count
		}

		if cursor == 0 {
			break
		}
	}

	return results, nil
}
