package service

import (
	"context"
	"log"
	v "root/genprotos/election"
	pb "root/genprotos/vote"
	"root/storage/postgres"
)

type VoteService struct {
	stg *postgres.VoteStorage
	pb.UnimplementedVoteServiceServer
}

func NewVoteService(stg *postgres.VoteStorage) *VoteService {
	return &VoteService{stg: stg}
}

func (c *VoteService) CreateVote(ctx context.Context, Vote *pb.Vote) (*v.Void, error) {
	v, err := c.stg.CreateVote(Vote)
	if err != nil {
		log.Print(err)
	}
	return v, err
}

func (c *VoteService) GetAllVotes(ctx context.Context, v *v.Void) (*pb.GetAllVote, error) {
	Votes, err := c.stg.GetAllVote(v)
	if err != nil {
		log.Print(err)
	}

	return Votes, err
}

func (c *VoteService) GetByIdVote(ctx context.Context, id *v.ById) (*pb.Vote, error) {
	prod, err := c.stg.GetByIdVote(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *VoteService) UpdateVote(ctx context.Context, Vote *pb.Vote) (*v.Void, error) {
	v, err := c.stg.UpdateVote(Vote)
	if err != nil {
		log.Print(err)
	}

	return v, err
}

func (c *VoteService) DeleteVote(ctx context.Context, id *v.ById) (*v.Void, error) {
	v, err := c.stg.DeleteVote(id)
	if err != nil {
		log.Print(err)
	}

	return v, err
}
