package service

import (
	"context"
	"log"
	pb "root/genprotos/election"
	"root/storage/postgres"
)

type ElectionService struct {
	stg *postgres.ElectionStorage
	pb.UnimplementedElectionServiceServer
}

func NewElectionService(stg *postgres.ElectionStorage) *ElectionService {
	return &ElectionService{stg: stg}
}

func (c *ElectionService) CreateElection(ctx context.Context, Election *pb.Election) (*pb.Void, error) {
	v, err := c.stg.CreateElection(Election)
	if err != nil {
		log.Print(err)
	}
	return v, err
}

func (c *ElectionService) GetAllElections(ctx context.Context, v *pb.Void) (*pb.GetAllElection, error) {
	Elections, err := c.stg.GetAllElection(v)
	if err != nil {
		log.Print(err)
	}

	return Elections, err
}

func (c *ElectionService) GetByIdElection(ctx context.Context, id *pb.ById) (*pb.Election, error) {
	prod, err := c.stg.GetByIdElection(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *ElectionService) UpdateElection(ctx context.Context, Election *pb.Election) (*pb.Void, error) {
	v, err := c.stg.UpdateElection(Election)
	if err != nil {
		log.Print(err)
	}

	return v, err
}

func (c *ElectionService) DeleteElection(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	v, err := c.stg.DeleteElection(id)
	if err != nil {
		log.Print(err)
	}

	return v, err
}
