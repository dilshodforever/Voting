package postgres

import (
	"database/sql"
	cn "root/genprotos"
	"github.com/google/uuid"
)

type CandidateStorage struct {
	db *sql.DB
}

func NewCandidateStorage(db *sql.DB) *CandidateStorage {
	return &CandidateStorage{db: db}
}

func (p *CandidateStorage) CreateCandidate(cn *cn.Candidate) (*cn.Void, error) {

	id := uuid.NewString()
	query := `
		INSERT INTO candidate (id, election_id, party_id, public_id, date)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := p.db.Exec(query, id, cn.Election.Id, cn.Public.Id, cn.Party.Id, cn.Date)
	return nil, err
}

func (p *CandidateStorage) GetByIdCandidate(id *cn.ById) (*cn.Candidate, error) {
		query := `
			SELECT id, election_id, party_id, public_id, date
			FROM candidate
			WHERE id = $1
		`
		row := p.db.QueryRow(query, id.Id)

		var can cn.Candidate
		err := row.Scan(&can.Election.Id,
			&can.Party.Id,
			&can.Public.Id,
			&can.Date)
		if err != nil {
			return nil, err
		}

		return &can, nil
}

func (p *CandidateStorage) GetAllCandidate(*cn.Void) (*cn.GetAllCandidate, error) {
	cans := &cn.GetAllCandidate{}
	row, err := p.db.Query("select id, election_id, party_id, public_id, date from candidate")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		can := &cn.Candidate{}
		err = row.Scan(&can.Election.Id, &can.Party.Id, &can.Public.Id)
		if err != nil {
			return nil, err
		}
		cans.Candidates=append(cans.Candidates, can)
	}
	return cans, nil
}

func (p *CandidateStorage) UpdateCandidate(cn *cn.Candidate) (*cn.Void, error) {
	query := `
		UPDATE candidate
		SET election_id=$1, party_id=$2, public_id=$3, date=$4
		WHERE id = $5
	`
	_, err := p.db.Exec(query,  cn.Election.Id, cn.Party.Id, cn.Public.Id)
	return nil, err
}

func (p *CandidateStorage) DeleteCandidate(id *cn.ById) (*cn.Void, error) {
	query := `
		delete from candidate where id = $1
	`
	_, err := p.db.Exec(query, id.Id)
	return nil, err
}
