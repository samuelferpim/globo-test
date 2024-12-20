package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"bbb-voting/internal/core/domain"
	"bbb-voting/internal/core/ports"

	"github.com/google/uuid"
)

type voteService struct {
	repo  ports.VoteRepository
	queue ports.QueueService
}

func NewVoteService(repo ports.VoteRepository, queue ports.QueueService) ports.VoteService {
	return &voteService{
		repo:  repo,
		queue: queue,
	}
}

func (uc *voteService) CastVote(ctx context.Context, vote *domain.Vote) error {
	log.Printf("Starting to cast vote: %+v", vote)

	if vote.ID == "" {
		vote.ID = uuid.New().String()
	}
	if vote.Timestamp.IsZero() {
		vote.Timestamp = time.Now()
	}

	voteRedis := domain.VoteRedis{
		ID:            vote.ID,
		ParticipantID: vote.ParticipantID,
		Timestamp:     vote.Timestamp,
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Storing vote in Redis: %+v", voteRedis)
		if err := uc.repo.StoreInRedis(ctx, &voteRedis); err != nil {
			log.Printf("Error storing vote in Redis: %v", err)
			errChan <- fmt.Errorf("failed to store vote in Redis: %w", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		voteJSON, err := json.Marshal(vote)
		if err != nil {
			log.Printf("Error marshalling vote: %v", err)
			errChan <- fmt.Errorf("failed to marshal vote: %w", err)
			return
		}
		log.Printf("Publishing vote to queue: %s", string(voteJSON))
		if err := uc.queue.Publish(ctx, "votes", voteJSON); err != nil {
			log.Printf("Error publishing vote to queue: %v", err)
			errChan <- fmt.Errorf("failed to publish vote to queue: %w", err)
		}
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			log.Printf("Error in CastVote: %v", err)
			return err
		}
	}

	log.Printf("Vote cast successfully: %s", vote.ID)
	return nil
}

func (s *voteService) GetDetailedResults(ctx context.Context) (domain.DetailedResults, error) {
	results, err := s.repo.GetResultsByParticipant(ctx)
	if err != nil {
		return domain.DetailedResults{}, err
	}

	totalVotes, err := s.repo.GetTotalVotes(ctx)
	if err != nil {
		return domain.DetailedResults{}, err
	}

	votesByHour, err := s.repo.GetVotesByHour(ctx)
	if err != nil {
		return domain.DetailedResults{}, err
	}

	participantResults := make([]domain.ParticipantResult, 0, len(results))
	for participantID, voteCount := range results {
		participantResults = append(participantResults, domain.ParticipantResult{
			ParticipantID: participantID,
			VoteCount:     voteCount,
		})
	}

	sort.Slice(participantResults, func(i, j int) bool {
		return participantResults[i].VoteCount > participantResults[j].VoteCount
	})

	return domain.DetailedResults{
		TotalVotes:         totalVotes,
		ParticipantResults: participantResults,
		VotesByHour:        votesByHour,
	}, nil
}

func (uc *voteService) GetTotalVotes(ctx context.Context) (int, error) {
	totalVotes, err := uc.repo.GetTotalVotes(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get total votes: %w", err)
	}

	return totalVotes, nil
}

func (uc *voteService) GetVotesByHour(ctx context.Context) (map[string]int, error) {
	votesByHour, err := uc.repo.GetVotesByHour(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get votes by hour: %w", err)
	}

	return votesByHour, nil
}
