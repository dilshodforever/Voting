package service

import (
	"context"
	"log"
	pb "root/genprotos"
	"root/storage"
)

type CandidateService struct {
	stg storage.InitRoot
	pb.UnimplementedCandidateServiceServer
}

func NewCandidateService(stg storage.InitRoot) *CandidateService {
	return &CandidateService{stg: stg}
}

func (c *CandidateService) CreateCandidate(ctx context.Context, Candidate *pb.Candidate) (*pb.Void, error) {
	pb, err := c.stg.Candidate().CreateCandidate(Candidate)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *CandidateService) GetAllCandidates(ctx context.Context, pb *pb.Void) (*pb.GetAllCandidate, error) {
	Candidates, err := c.stg.Candidate().GetAllCandidate(pb)
	if err != nil {
		log.Print(err)
	}

	return Candidates, err
}

func (c *CandidateService) GetByIdCandidate(ctx context.Context, id *pb.ById) (*pb.Candidate, error) {
	prod, err := c.stg.Candidate().GetByIdCandidate(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *CandidateService) UpdateCandidate(ctx context.Context, Candidate *pb.Candidate) (*pb.Void, error) {
	pb, err := c.stg.Candidate().UpdateCandidate(Candidate)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *CandidateService) DeleteCandidate(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Candidate().DeleteCandidate(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}
