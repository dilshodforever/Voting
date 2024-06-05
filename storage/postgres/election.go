package postgres

import (
	"database/sql"
	el "root/genprotos/election"
	"github.com/google/uuid"
	v "root/genprotos/election"
)

type ElectionStorage struct {
	db *sql.DB
}

func NewElectionStorage(db *sql.DB) *ElectionStorage {
	return &ElectionStorage{db: db}
}
            
func (bc *ElectionStorage) CreateElection(el *el.Election ) (*el.Void,error){
	id:=uuid.NewString()
	_, err:=bc.db.Exec(`insert into election(id, name, date) 
						values($1, $2)`,
						id, el.Name, el.Date)
	return nil,err
}


func (p *ElectionStorage) GetByIdElection(id *v.ById) (*el.Election, error) {
	query := `
		SELECT id, name, date
		FROM election
		WHERE id = $1
	`
	row := p.db.QueryRow(query, id.Id)

	var elec el.Election
	err := row.Scan(&elec.Id,
		&elec.Name,
		&elec.Date)
	if err != nil {
		return nil, err
	}

	return &elec, nil
}

func (p *ElectionStorage) GetAllElection(*v.Void) (*el.GetAllElection, error) {
elecs := &el.GetAllElection{}
row, err := p.db.Query("select id, name, date from election")
if err != nil {
	return nil, err
}
for row.Next() {
	elec := &el.Election{}
	err = row.Scan(&elec.Id, &elec.Name, &elec.Date)
	if err != nil {
		return nil, err
	}
	elecs.Elections=append(elecs.Elections, elec)
}
return elecs, nil
}

func (p *ElectionStorage) UpdateElection(el *el.Election) (*v.Void, error) {
query := `
	UPDATE election
	SET  name=$1, date=$2
	WHERE id = $3
`
_, err := p.db.Exec(query,  el.Name, el.Date)
return nil, err
}

func (p *ElectionStorage) DeleteElection(id *v.ById) (*v.Void, error) {
query := `
	delete from election  where id = $1
`
_, err := p.db.Exec(query, id.Id)
return nil, err
}
