package postgres

import (
	"database/sql"
	v "root/genprotos/election"
	pb "root/genprotos/public"

	"github.com/google/uuid"
)

type PublicStorage struct {
	db *sql.DB
}

func NewPublicStorage(db *sql.DB) *PublicStorage {
	return &PublicStorage{db: db}
}

func (p *PublicStorage) CreatePublic(pub *pb.Public) (*v.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO public_ (id, election_id, public_id)
		VALUES ($1, $2, $3)
	`
	_, err := p.db.Exec(query, id, p.Election.Id, p.Public.Id)
	return nil, err
}

func (p *PublicStorage) GetByIdPublic(id *v.ById) (*pb.Public, error) {
		query := `
			SELECT id, name
			FROM public_, election_id, public_id
			WHERE id = $1
		`
		row := p.db.QueryRow(query, id.Id)

		var pub pb.Public
		err := row.Scan(&p.Id,
			&p.Election.Id,
			&p.Public.Id)
		if err != nil {
			return nil, err
		}

		return &p, nil
}

func (p *PublicStorage) GetAllPublic(*v.Void) (*pb.GetAllPublic, error) {
	ps := &pb.GetAllPublic{}
	row, err := p.db.Query("select id, election_id, public_id from public_")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		pub := &p.CreatePublic{}
		err = row.Scan(&p.Id, &p.Election.Id, &p.Public.Id)
		if err != nil {
			return nil, err
		}
		ps.Publics = append(ps.Publics, p)
	}
	return users, nil
}

func (pub *PublicStorage) UpdatePublic(p *pb.Public) (*v.Void, error) {
	query := `
		UPDATE public_
		SET election_id = $2, public_id = $3 
		WHERE id = $1 
	`
	_, err := p.db.Exec(query, p.Id, p.Election.Id, p.Public.Id)
	return nil, err
}

func (p *PublicStorage) Delete(id *v.ById) (*v.Void, error) {
	query := `
		delete from public_ where id = $1
	`
	_, err := p.db.Exec(query, id.Id)
	return nil, err
}
