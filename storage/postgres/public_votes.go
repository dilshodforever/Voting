package postgres

import (
	"database/sql"
	pb "root/genprotos"

	"github.com/google/uuid"
)

type PublicVoteStorage struct {
	db *sql.DB
}

func NewPublicVoteStorage(db *sql.DB) *PublicVoteStorage {
	return &PublicVoteStorage{db: db}
}

func (p *PublicVoteStorage) CreatePublicVote(pVote *pb.PublicVote) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO public_vote (id, election_id, public_id)
		VALUES ($1, $2, $3)
	`
	_, err := p.db.Exec(query, id, pVote.Election.Id, pVote.Public.Id)
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
		err := row.Scan(&pVote.Id,
			&pVote.Election.Id,
			&pVote.Public.Id)
		if err != nil {
			return nil, err
		}

		return &pVote, nil
}

func (p *PublicVoteStorage) GetAllPublicVote(_ *pb.Void) (*pb.GetAllPublicVote, error) {
	pVotes := &pb.GetAllPublicVote{}
	row, err := p.db.Query("select id, election_id, public_id from public_vote")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		pVote := &pb.PublicVote{}
		err = row.Scan(&pVote.Id, &pVote.Election.Id, &pVote.Public.Id)
		if err != nil {
			return nil, err
		}
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
		delete from public_vote where id = $1
	`
	_, err := p.db.Exec(query, id.Id)
	return nil, err
}
