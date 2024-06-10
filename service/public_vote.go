package service

import (
	"context"
	"log"
	pb "root/genprotos"
	"root/storage/postgres"
)

type PublicVoteService struct {
	stg *postgres.PublicVoteStorage
	pb.UnimplementedPublicVoteServiceServer
}

func NewPublicVoteService(stg *postgres.PublicVoteStorage) *PublicVoteService {
	return &PublicVoteService{stg: stg}
}

func (c *PublicVoteService) CreatePublicVote(ctx context.Context, publicVote *pb.PublicVote) (*pb.Void, error) {
	pb, err := c.stg.CreatePublicVote(publicVote)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *PublicVoteService) GetAllPublicVotes(ctx context.Context, pb *pb.PublicVote) (*pb.GetAllPublicVote, error) {
	publicVotes, err := c.stg.GetAllPublicVote(pb)
	if err != nil {
		log.Print(err)
	}

	return publicVotes, err
}

func (c *PublicVoteService) GetByIdPublicVote(ctx context.Context, id *pb.ById) (*pb.PublicVote, error) {
	prod, err := c.stg.GetByIdPublicVote(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *PublicVoteService) UpdatePublicVote(ctx context.Context, publicVote *pb.PublicVote) (*pb.Void, error) {
	pb, err := c.stg.UpdatePublicVote(publicVote)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *PublicVoteService) DeletePublicVote(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.DeletePublicVote(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}
