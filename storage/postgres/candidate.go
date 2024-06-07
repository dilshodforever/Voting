package postgres

import (
	"database/sql"
	cn "root/genprotos"
	"time"

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
		INSERT INTO candidate (id, election_id, party_id, public_id)
		VALUES ($1, $2, $3, $4)
	`

	_, err := p.db.Exec(query, id, cn.Election.Id, cn.Public.Id, cn.Party.Id)
	return nil, err
}

func (p *CandidateStorage) GetByIdCandidate(id *cn.ById) (*cn.Candidate, error) {
	query := `
			SELECT id, election_id, party_id, public_id, date
			FROM candidate
			WHERE id = $1
		`
	row := p.db.QueryRow(query, id.Id)

	can := &cn.Candidate{}
	var eid, pid, puid string
	err := row.Scan(&eid, &pid, &puid)
	if err != nil {
		return nil, err
	}
	can.Election.Id = eid
	can.Party.Id = pid
	can.Public.Id = puid
	return can, nil
}

func (p *CandidateStorage) GetAllCandidate(*cn.Void) (*cn.GetAllCandidate, error) {
	cans := &cn.GetAllCandidate{}
	row, err := p.db.Query(`select  election_id, party_id, public_id from candidate
						     where delated_at=0`)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		can := &cn.Candidate{}
		var eid, pid, puid string
		err = row.Scan(&eid, &pid, &puid)
		if err != nil {
			return nil, err
		}
		can.Election.Id = eid
		can.Party.Id = pid
		can.Public.Id = puid
		cans.Candidates = append(cans.Candidates, can)
	}
	return cans, nil
}

func (p *CandidateStorage) UpdateCandidate(cn *cn.Candidate) (*cn.Void, error) {
	query := `
		UPDATE candidate
		SET election_id=$1, party_id=$2, public_id=$3
		WHERE id = $4
	`
	_, err := p.db.Exec(query, cn.Election.Id, cn.Party.Id, cn.Public.Id)
	return nil, err
}

func (p *CandidateStorage) DeleteCandidate(id *cn.ById) (*cn.Void, error) {
	query := `
		update   candidate set delated_at=$1 where id = $2
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}
