package service

import (
	"context"
	"log"
	v "root/genprotos/election"
	pb "root/genprotos/candidate"
	"root/storage/postgres"
)

type CandidateService struct {
	stg *postgres.CandidateStorage
	pb.UnimplementedCandidateServiceServer
}

func NewCandidateService(stg *postgres.CandidateStorage) *CandidateService {
	return &CandidateService{stg: stg}
}

func (c *CandidateService) CreateCandidate(ctx context.Context, Candidate *pb.Candidate) (*v.Void, error) {
	v, err := c.stg.CreateCandidate(Candidate)
	if err != nil {
		log.Print(err)
	}
	return v, err
}

func (c *CandidateService) GetAllCandidates(ctx context.Context, v *v.Void) (*pb.GetAllCandidate, error) {
	Candidates, err := c.stg.GetAllCandidate(v)
	if err != nil {
		log.Print(err)
	}

	return Candidates, err
}

func (c *CandidateService) GetByIdCandidate(ctx context.Context, id *v.ById) (*pb.Candidate, error) {
	prod, err := c.stg.GetByIdCandidate(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *CandidateService) UpdateCandidate(ctx context.Context, Candidate *pb.Candidate) (*v.Void, error) {
	v, err := c.stg.UpdateCandidate(Candidate)
	if err != nil {
		log.Print(err)
	}

	return v, err
}

func (c *CandidateService) DeleteCandidate(ctx context.Context, id *v.ById) (*v.Void, error) {
	v, err := c.stg.DeleteCandidate(id)
	if err != nil {
		log.Print(err)
	}

	return v, err
}
