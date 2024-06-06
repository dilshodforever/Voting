package service

import (
	"database/sql"
	"root/storage/postgres"
)

type Service struct {
	candidateService  *CandidateService
	electionService   *ElectionService
	publicVoteService *PublicVoteService
}

func InitServices(db *sql.DB) *Service {
	candidate := NewCandidateService(postgres.NewCandidateStorage(db))
	election := NewElectionService(postgres.NewElectionStorage(db))
	publicVote := NewPublicVoteService(postgres.NewPublicVoteStorage(db))
	return &Service{candidate, election, publicVote}
}
