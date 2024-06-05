package postgres

import (
	"database/sql"
	"fmt"
	"root/config"
	"root/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db      *sql.DB

	Candidates storage.Candidate
	Elections storage.Election
	PublicVotes storage.PublicVote

	Votes storage.Vote
}


func NewPostgresStorage() (storage.InitRoor, error) {
	config:=config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
						config.PostgresUser, config.PostgresPassword, 
						config.PostgresHost, config.PostgresPort, 
						config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	 return &Storage{Db: db}, err

}


func (s *Storage) Candidate() storage.Candidate {
	if s.Candidates == nil {
		s.Candidates =&CandidateStorage{s.Db}
	}
	return s.Candidates
}



func (s *Storage) Election() storage.Election {
	if s.Elections == nil {
		s.Elections = &ElectionStorage{s.Db}
	}
	return s.Elections
}

func (s *Storage) PublicVote() storage.PublicVote {
	if s.PublicVotes == nil {
		s.PublicVotes= &PublicVoteStorage{s.Db}
	}
	return s.PublicVotes
}

func (s *Storage) Vote() storage.Vote {
	if s.Votes == nil {
		s.Votes= &VoteStorage{s.Db}
	}
	return s.Votes
}

