package postgres

import (
	"database/sql"
	v "root/genprotos/election"
	pb "root/genprotos/vote"

	"github.com/google/uuid"
)

type VoteStorage struct {
	db *sql.DB
}

func NewVoteStorage(db *sql.DB) *VoteStorage {
	return &VoteStorage{db: db}
}

func (v *VoteStorage) CreateVote(vote *pb.Vote) (*v.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO vote (id, candidate_id, data)
		VALUES ($1, $2, $3)
	`
	_, err := v.db.Exec(query, id, vote.Candidate.Id, vote.Date)
	return nil, err
}

func (v *VoteStorage) GetByIdVote(id *v.ById) (*pb.Vote, error) {
	query := `
			SELECT id, candidate_id, data
			FROM vote
			WHERE id = $1
		`
	row := v.db.QueryRow(query, id.Id)

	var vote pb.Vote
	err := row.Scan(&vote.Id,
		&vote.Candidate.Id,
		&vote.Date)
	if err != nil {
		return nil, err
	}

	return &vote, nil
}

func (v *VoteStorage) GetAllVote(_ *v.Void) (*pb.GetAllVote, error) {
	votes := &pb.GetAllVote{}
	row, err := v.db.Query("select id, candidate_id, data from vote")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		vote := &pb.Vote{}
		err = row.Scan(&vote.Id, &vote.Candidate.Id, &vote.Date)
		if err != nil {
			return nil, err
		}
		votes.Votes = append(votes.Votes, vote)
	}
	return votes, nil
}

func (p *VoteStorage) UpdateVote(vote *pb.Vote) (*v.Void, error) {
	query := `
		UPDATE Vote_
		SET candidate_id = $2, data = $3
		WHERE id = $1 
	`
	_, err := p.db.Exec(query, vote.Id, vote.Candidate.Id, vote.Date)
	return nil, err
}

func (v *VoteStorage) DeleteVote(id *v.ById) (*v.Void, error) {
	query := `
		delete from vote where id = $1
	`
	_, err := v.db.Exec(query, id.Id)
	return nil, err
}
