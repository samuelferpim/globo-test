package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"bbb-voting/internal/core/domain"
	"bbb-voting/internal/core/ports"
	"bbb-voting/internal/repository"
)

type voteRepository struct {
	redis repository.RedisClient
}

func NewVoteRepository(redis repository.RedisClient) ports.VoteRepository {
	return &voteRepository{
		redis: redis,
	}
}

func (r *voteRepository) StoreInRedis(ctx context.Context, vote *domain.VoteRedis) error {
	pipe := r.redis.Pipeline()

	voteKey := fmt.Sprintf("vote:%s", vote.ID)
	pipe.HMSet(ctx, voteKey, map[string]interface{}{
		"id":             vote.ID,
		"participant_id": vote.ParticipantID,
		"timestamp":      vote.Timestamp.Unix(),
	})

	participantCountKey := fmt.Sprintf("count:participant:%s", vote.ParticipantID)
	pipe.Incr(ctx, participantCountKey)

	pipe.Incr(ctx, "count:total")

	hourKey := vote.Timestamp.Format("2006010215")
	pipe.HIncrBy(ctx, "votes:by:hour", hourKey, 1)

	return pipe.Exec(ctx)
}

func (r *voteRepository) GetResultsByParticipant(ctx context.Context) (map[string]int, error) {
	var cursor uint64
	results := make(map[string]int)

	for {
		keys, nextCursor, err := r.redis.Scan(ctx, cursor, "count:participant:*", 10).Result()
		if err != nil {
			return nil, err
		}

		for _, key := range keys {
			participantID := strings.TrimPrefix(key, "count:participant:")
			countStr, err := r.redis.Get(ctx, key).Result()
			if err != nil {
				return nil, err
			}
			count, err := strconv.Atoi(countStr)
			if err != nil {
				return nil, err
			}
			results[participantID] = count
		}

		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	return results, nil
}

func (r *voteRepository) GetTotalVotes(ctx context.Context) (int, error) {
	countStr, err := r.redis.Get(ctx, "count:total").Result()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(countStr)
}

func (r *voteRepository) GetVotesByHour(ctx context.Context) (map[string]int, error) {
	results, err := r.redis.HGetAll(ctx, "votes:by:hour").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get votes by hour: %w", err)
	}

	votesByHour := make(map[string]int, len(results))
	for hour, countStr := range results {
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return nil, fmt.Errorf("failed to convert count to int for hour %s: %w", hour, err)
		}
		votesByHour[hour] = count
	}

	return votesByHour, nil
}
