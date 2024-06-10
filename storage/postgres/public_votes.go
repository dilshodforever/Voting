package postgres

import (
	"database/sql"
	"fmt"
	pb "root/genprotos"
	"time"

	"github.com/google/uuid"
)

type PublicVoteStorage struct {
	db *sql.DB
}

func NewPublicVoteStorage(db *sql.DB) *PublicVoteStorage {
	return &PublicVoteStorage{db: db}
}

func (p *PublicVoteStorage) CreatePublicVote(pVote *pb.PublicVote) (*pb.Void, error) {
	t, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer t.Commit()
	id := uuid.NewString()
	query := `
		INSERT INTO public_vote (id, election_id, public_id)
		VALUES ($1, $2, $3)
	`
	_, err = t.Exec(query, id, pVote.Election.Id, pVote.Public.Id)
	if err != nil {
		return nil, err
	}
	id = uuid.NewString()
	query = `
		INSERT INTO vote (id, candidate_id)
		VALUES ($1, $2)`
	_, err = t.Exec(query, id, pVote.Candidate.Id)

	return nil, err
}

func (p *PublicVoteStorage) GetByIdPublicVote(id *pb.ById) (*pb.PublicVote, error) {
	query := `
			SELECT id, election_id, public_id
			FROM public_vote
			WHERE id = $1
		`
	row := p.db.QueryRow(query, id.Id)

	var pVote pb.PublicVote
	var eid, pid string
	err := row.Scan(&pVote.Id,
		&eid,
		&pid)
	if err != nil {
		return nil, err
	}
	pVote.Election.Id = eid
	pVote.Public.Id = pid

	return &pVote, nil
}

func (p *PublicVoteStorage) GetAllPublicVote(pu *pb.PublicVote) (*pb.GetAllPublicVote, error) {
	pVotes := &pb.GetAllPublicVote{}
	count := 1
	query := `select id, election_id, public_id from public_vote where delated_at=0`
	var arr []interface{}
	if len(pu.Election.Id) > 0 {
		query += fmt.Sprintf(` and election_id=$%d`, count)
		arr = append(arr, pu.Election.Id)
	}
	if len(pu.Public.Id) > 0 {
		query += fmt.Sprintf(` and public_id=$%d`, count)
		arr = append(arr, pu.Public.Id)
	}
	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		pVote := &pb.PublicVote{}
		var eid, pid string
		err = row.Scan(&pVote.Id, &eid, &pid)
		if err != nil {
			return nil, err
		}
		pVote.Election.Id = eid
		pVote.Public.Id = pid
		pVotes.PublicVotes = append(pVotes.PublicVotes, pVote)
	}
	return pVotes, nil
}

func (p *PublicVoteStorage) UpdatePublicVote(pVote *pb.PublicVote) (*pb.Void, error) {
	query := `
		UPDATE public_vote
		SET election_id = $2, public_id = $3 
		WHERE id = $1 
	`
	_, err := p.db.Exec(query, pVote.Id, pVote.Election.Id, pVote.Public.Id)
	return nil, err
}

func (p *PublicVoteStorage) DeletePublicVote(id *pb.ById) (*pb.Void, error) {
	query := `
		update  public_vote set delated_at=$1 where id = $2
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}
