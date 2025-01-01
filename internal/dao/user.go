package dao

import (
	"database/sql"

	"github.com/print-Eruki/CIIC4060-chatbot/internal/model"
)

type UserDAO struct {
	DB *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{DB: db}
}

// Creates a new user and modifies the param in-memory
func (dao *UserDAO) CreateUser(newUser *model.User) error {
	query := `
	INSERT INTO public.user
		(username, password)
	VALUES 
		($1, $2)
	RETURNING
		uid, username, password, created_at
	`

	err := dao.DB.QueryRow(query, newUser.Username, newUser.Password).Scan(
		&newUser.Uid,
		&newUser.Username,
		&newUser.Password,
		&newUser.Created_at,
	)

	return err
}

func (dao *UserDAO) GetUser(username string) (*model.User, error) {
	query := `
	SELECT 
		uid, username, password, created_at
	FROM 
		public.user
	WHERE 
		username = $1;
	`
	var user model.User
	err := dao.DB.QueryRow(query, username).Scan(
		&user.Uid,
		&user.Username,
		&user.Password,
		&user.Created_at,
	)
	if err != nil {
		return &model.User{}, err
	}

	return &user, nil
}
