package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// service ini membutuhkan repository

type ServiceUser interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
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




func (s *serviceUser) LoginUser(input LoginInput) (User, error) {
	// mapping struct input ke struct user
	// simpan struct user melalui repository

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user not found")
	} 

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}


func (s *serviceUser) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	// Mengambil email dari input
	email := input.Email

	// Mencari pengguna berdasarkan email dari repository
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		// Jika terjadi kesalahan saat mencari email, kembalikan false dan error
		return false, err
	}

	// Jika user dengan email yang diberikan tidak ditemukan
	if user.ID == 0 {
		// Maka email tidak tersedia, kembalikan true dan tidak ada error
		return true, nil
	}	

	// Jika email telah digunakan, kembalikan false dan tidak ada error
	return false, nil
}

