package user

import (
	"golang.org/x/crypto/bcrypt"
)

// service ini membutuhkan repository

type ServiceUser interface {
	RegisterUser(input RegisterUserInput) (User, error)
}


type serviceUser struct {
   	repository RepositoryUser
}


func NewServiceUser(repository RepositoryUser) *serviceUser {
	return &serviceUser{repository: repository}
}


func (s *serviceUser) RegisterUser(input RegisterUserInput) (User, error) {
	// mapping
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

    passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	// hard core
	user.Role = "user"


    newUser, err := s.repository.Save(user)
	if err != nil {
		return user, nil
	}

	return newUser, nil
}

