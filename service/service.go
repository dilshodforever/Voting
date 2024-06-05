package service

import (
	"database/sql"
	"root/storage/postgres"
)

type Service struct {
	candidateService   *CandidateService
	electionService   *ElectionService
	partyService   *PartyService
	publicVoteService   *PublicVoteService
	publicService   *PublicService
	voteService   *VoteService
}

func InitServices(db *sql.DB) *Service {
	candidate := NewCandidateService(postgres.NewCandidateStorage(db))
	election := NewElectionService(postgres.NewElectionStorage(db))
	party := NewPartyService(postgres.NewPartyStorage(db))
	publicVote := NewPublicVoteService(postgres.NewPublicVoteStorage(db))
	public := NewPublicService(postgres.NewPublicStorage(db))
	vote := NewVoteService(postgres.NewVoteStorage(db))

	return &Service{candidate, election, party, publicVote, public, vote}
}