package service

import (
	"context"
	"errors"

	"bbb-voting/internal/core/domain"
	"bbb-voting/internal/core/ports"
	. "bbb-voting/tests/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("VoteService", func() {
	var (
		mockCtrl    *gomock.Controller
		mockRepo    *MockVoteRepository
		mockQueue   *MockQueueService
		voteService ports.VoteService
		ctx         context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = NewMockVoteRepository(mockCtrl)
		mockQueue = NewMockQueueService(mockCtrl)
		voteService = NewVoteService(VoteServiceOptions{
			Repo:  mockRepo,
			Queue: mockQueue,
		})
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("CastVote", func() {
		It("should successfully cast a vote", func() {
			vote := &domain.Vote{
				ParticipantID: "participant1",
			}

			mockRepo.EXPECT().
				StoreInRedis(gomock.Any(), gomock.Any()).
				Return(nil)

			mockQueue.EXPECT().
				Publish(gomock.Any(), "votes", gomock.Any()).
				Return(nil)

			err := voteService.CastVote(ctx, vote)

			Expect(err).To(BeNil())
			Expect(vote.ID).NotTo(BeEmpty())
			Expect(vote.Timestamp).NotTo(BeZero())
		})

		It("should return an error if Redis storage fails", func() {
			vote := &domain.Vote{
				ParticipantID: "participant1",
			}

			mockRepo.EXPECT().
				StoreInRedis(gomock.Any(), gomock.Any()).
				Return(errors.New("Redis error"))

			mockQueue.EXPECT().
				Publish(gomock.Any(), "votes", gomock.Any()).
				Return(nil)

			err := voteService.CastVote(ctx, vote)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to store vote in Redis"))
		})

		It("should return an error if queue publish fails", func() {
			vote := &domain.Vote{
				ParticipantID: "participant1",
			}

			mockRepo.EXPECT().
				StoreInRedis(gomock.Any(), gomock.Any()).
				Return(nil)

			mockQueue.EXPECT().
				Publish(gomock.Any(), "votes", gomock.Any()).
				Return(errors.New("Queue error"))

			err := voteService.CastVote(ctx, vote)

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to publish vote to queue"))
		})
	})

	Describe("GetDetailedResults", func() {
		It("should return detailed results", func() {
			mockResults := map[string]int{
				"participant1": 100,
				"participant2": 200,
			}
			mockTotalVotes := 300
			mockVotesByHour := map[string]int{
				"2023-05-01 10:00": 150,
				"2023-05-01 11:00": 150,
			}

			mockRepo.EXPECT().
				GetResultsByParticipant(gomock.Any()).
				Return(mockResults, nil)

			mockRepo.EXPECT().
				GetTotalVotes(gomock.Any()).
				Return(mockTotalVotes, nil)

			mockRepo.EXPECT().
				GetVotesByHour(gomock.Any()).
				Return(mockVotesByHour, nil)

			results, err := voteService.GetDetailedResults(ctx)

			Expect(err).To(BeNil())
			Expect(results.TotalVotes).To(Equal(mockTotalVotes))
			Expect(results.ParticipantResults).To(HaveLen(2))
			Expect(results.VotesByHour).To(Equal(mockVotesByHour))
		})

		It("should return an error if GetResultsByParticipant fails", func() {
			mockRepo.EXPECT().
				GetResultsByParticipant(gomock.Any()).
				Return(nil, errors.New("Database error"))

			results, err := voteService.GetDetailedResults(ctx)

			Expect(err).To(HaveOccurred())
			Expect(results).To(Equal(domain.DetailedResults{}))
		})
	})

	Describe("GetTotalVotes", func() {
		It("should return total votes", func() {
			mockTotalVotes := 500

			mockRepo.EXPECT().
				GetTotalVotes(gomock.Any()).
				Return(mockTotalVotes, nil)

			totalVotes, err := voteService.GetTotalVotes(ctx)

			Expect(err).To(BeNil())
			Expect(totalVotes).To(Equal(mockTotalVotes))
		})

		It("should return an error if GetTotalVotes fails", func() {
			mockRepo.EXPECT().
				GetTotalVotes(gomock.Any()).
				Return(0, errors.New("Database error"))

			totalVotes, err := voteService.GetTotalVotes(ctx)

			Expect(err).To(HaveOccurred())
			Expect(totalVotes).To(Equal(0))
		})
	})

	Describe("GetVotesByHour", func() {
		It("should return votes by hour", func() {
			mockVotesByHour := map[string]int{
				"2023-05-01 10:00": 150,
				"2023-05-01 11:00": 150,
			}

			mockRepo.EXPECT().
				GetVotesByHour(gomock.Any()).
				Return(mockVotesByHour, nil)

			votesByHour, err := voteService.GetVotesByHour(ctx)

			Expect(err).To(BeNil())
			Expect(votesByHour).To(Equal(mockVotesByHour))
		})

		It("should return an error if GetVotesByHour fails", func() {
			mockRepo.EXPECT().
				GetVotesByHour(gomock.Any()).
				Return(nil, errors.New("Database error"))

			votesByHour, err := voteService.GetVotesByHour(ctx)

			Expect(err).To(HaveOccurred())
			Expect(votesByHour).To(BeNil())
		})
	})
})
