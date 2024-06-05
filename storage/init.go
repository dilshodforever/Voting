package storage

import (
	pb "root/genprotos"
)

type InitRoor interface {
	Candidate() Candidate
	Election() Election
	PublicVote() PublicVote

	Vote() Vote
}

type Candidate interface {
	CreateCandidate(cn *pb.Candidate) (*pb.Void, error)
	GetByIdCandidate(id *pb.ById) (*pb.Candidate, error)
	GetAllCandidate(*pb.Void) (*pb.GetAllCandidate, error)
	UpdateCandidate(cn *pb.Candidate) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
}

type Election interface {
	CreateElection(el *pb.Election) (*pb.Void, error)
	GetByIdElection(el *pb.ById) (*pb.Election, error)
	GetAllElection(*pb.Void) (*pb.GetAllElection, error)
	UpdateElection(el *pb.Election) (*pb.Void, error)
	DeleteElection(id *pb.ById) (*pb.Void, error)
}

type PublicVote interface {
	CreatePublicVote(pVote *pb.PublicVote) (*pb.Void, error)
	GetByIdPublicVote(id *pb.ById) (*pb.PublicVote, error)
	GetAllPublicVote(_ *pb.Void) (*pb.GetAllPublicVote, error)
	UpdatePublicVote(pVote *pb.PublicVote) (*pb.Void, error)
	DeletePublicVote(id *pb.ById) (*pb.Void, error)
}

type Vote interface {
	CreateVote(vote *pb.Vote) (*pb.Void, error)
	GetByIdVote(id *pb.ById) (*pb.Vote, error)
	GetAllVote(_ *pb.Void) (*pb.GetAllVote, error)
	UpdateVote(vote *pb.Vote) (*pb.Void, error)
	DeleteVote(id *pb.ById) (*pb.Void, error)
}
