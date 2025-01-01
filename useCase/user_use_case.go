package useCase

import (
	"go-api/model"
	"go-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repo,
	}
}

func (uc *UserUseCase) GetUsers() ([]model.User, error) {
	return uc.repository.GetUsers()
}

func (uc *UserUseCase) GetUserByID(id int) (*model.User, error) {
	user, err := uc.repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) CreateUser(user model.User) (model.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return model.User{}, err
	}

	user.Password = string(hash)

	userID, err := uc.repository.CreateUser(user)

	if err != nil {
		return model.User{}, err
	}
	user.ID = userID
	return user, nil
}

func (uc *UserUseCase) Login(email string, password string) (*model.User, error) {
	user, err := uc.repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return user, nil
}
