package repository

import (
	"context"
	"database/sql"
	"test_tracker_backend/domain"
)

 type userRepository struct{
	db *sql.DB
 }

 // newuserRepository instantiates new data access layer for user
func NewUserRepository (db *sql.DB) domain.UserRepository{
	return  &userRepository{db : db}
}

func (r *userRepository) Create(ctx context.Context , user *domain.User) error{
	query := 
	`INSERT INTO users (_id, name, passwordhash , created_at)
	VALUES ($1 $2 $3 $4)`

	_,err := r.db.ExecContext(ctx,query,user.ID,user.Name, user.PasswordHash,user.CreatedAt)
	return err
}
func (r *userRepository) GetByName(ctx context.Context , name string) (*domain.User,error){
	query := `SELECT _id,name,passwordhash,created_at FROM users where name = $1;`

	row := r.db.QueryRowContext(ctx, query, name)
	
	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, err
	}
	return &user, nil
}

func (r * userRepository) GetByID(ctx context.Context, ID string) (*domain.User, error){

	query := `SELECT _id,name,passwordhash,created_at FROM users where name = $1;`

	row := r.db.QueryRowContext(ctx, query, ID)
	
	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, err
	}
	return &user, nil

}


