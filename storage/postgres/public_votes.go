package postgres

import (
	"database/sql"
	user "root/genprotos/user"
	"github.com/google/uuid"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}
func (u *UserStorage) CreateUser(user *user.CreateUser) (*user.UserREsponse,error) {
	id := uuid.NewString()
	query := `
		INSERT INTO Users (id, name, email)
		VALUES ($1, $2, $3)
	`
	_, err := u.db.Exec(query, id, user.Name,user.Email)
	return nil,err
}

func (u *UserStorage) GetAllUser(fl *user.Filter) (*user.GetAllUsers,error) {
	users:=&user.GetAllUsers{}
	row, err:=u.db.Query("select name, email from users where deleted_at=0")
	if err!=nil{
		return nil, err
	}
	for row.Next(){
		user:=&user.CreateUser{}
		err=row.Scan(&user.Name, &user.Email)
		if err!=nil{
			return nil, err
		}
		users.Users=append(users.Users, user)
	}
	return users, nil
}


