package postgres

import (
	"database/sql"
	election "root/genprotos/election"
	"github.com/google/uuid"
)

type ElectionStorage struct {
	Db *sql.DB
}

func NewElectionStorage(db *sql.DB) *ElectionStorage {
	return &ElectionStorage{Db: db}
}
            
func (bc *ElectionStorage) CreateElection(el *election.Election ) (*election.Void,error){
	id:=uuid.NewString()
	_, err:=bc.Db.Exec(`insert into election(id, name, date) 
						values($1, $2)`,
						id, el.Name, el.Date)
	return nil,err
}

