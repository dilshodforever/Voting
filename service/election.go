package service

import (
	"context"
	"log"
	pb "root/genprotos"
	"root/storage"
)

type ElectionService struct {
	stg storage.InitRoot
	pb.UnimplementedElectionServiceServer
}

func NewElectionService(stg storage.InitRoot) *ElectionService {
	return &ElectionService{stg: stg}
}

func (c *ElectionService) CreateElection(ctx context.Context, Election *pb.Election) (*pb.Void, error) {
	v, err := c.stg.Election().CreateElection(Election)
	if err != nil {
		log.Print(err)
	}
	return v, err
}

func (c *ElectionService) GetAllElections(ctx context.Context, v *pb.Election) (*pb.GetAllElection, error) {
	Elections, err := c.stg.Election().GetAllElection(v)
	if err != nil {
		log.Print(err)
	}

	return Elections, err
}

func (c *ElectionService) GetByIdElection(ctx context.Context, id *pb.ById) (*pb.Election, error) {
	prod, err := c.stg.Election().GetByIdElection(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *ElectionService) UpdateElection(ctx context.Context, Election *pb.Election) (*pb.Void, error) {
	v, err := c.stg.Election().UpdateElection(Election)
	if err != nil {
		log.Print(err)
	}

	return v, err
}

func (c *ElectionService) DeleteElection(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	v, err := c.stg.Election().DeleteElection(id)
	if err != nil {
		log.Print(err)
	}

	return v, err
}
