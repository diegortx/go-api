package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		connection: db,
	}
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {
	query := "SELECT id,name,email,password FROM users"
	rows, err := ur.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.User{}, err

	}

	var userList []model.User
	var userObj model.User

	for rows.Next() {
		err := rows.Scan(
			&userObj.ID,
			&userObj.Name,
			&userObj.Email,
			&userObj.Password,
		)
		if err != nil {
			fmt.Println(err)
			return []model.User{}, err
		}

		userList = append(userList, userObj)
	}

	rows.Close()

	return userList, nil
}

func (ur *UserRepository) GetUserByID(id int) (*model.User, error) {
	query, err := ur.connection.Prepare("SELECT id,name,email,password FROM users WHERE id=$1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user model.User

	err = query.QueryRow(id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &user, nil
}

func (ur *UserRepository) CreateUser(user model.User) (int, error) {
	var id int

	checkEmail, _ := ur.GetUserByEmail(user.Email)

	if checkEmail != nil {
		return 0, fmt.Errorf("email already exists")
	}

	query, err := ur.connection.Prepare("INSERT INTO users(name,email,password) VALUES($1,$2,$3) returning id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = query.QueryRow(user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	query, err := ur.connection.Prepare("SELECT id,name,email,password FROM users WHERE email =$1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user model.User

	err = query.QueryRow(email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &user, nil
}
