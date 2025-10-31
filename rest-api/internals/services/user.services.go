package services

import (
	"rest-api/internals/repository"
	"rest-api/x/interfacesx"
)

type UserService interface {
	CreateUserAccount(userRequest *interfacesx.UserRegistrationRequest) (*interfacesx.UserData, error)
	FetchUserAccount(userEmail string) (*interfacesx.UserData, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

// Method to create user
func (us *userService) CreateUserAccount(userRequest *interfacesx.UserRegistrationRequest) (*interfacesx.UserData, error) {
	userData, err := us.userRepo.CreateUserAccount(userRequest)
	if err != nil {
		return nil, err
	}

	return &interfacesx.UserData{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		Username:  userData.Username,
		UserRole:  userData.UserRole,
		CreatedAt: userData.CreatedAt,
	}, nil
}

// Method to fetch user details
func (us *userService) FetchUserAccount(userEmail string) (*interfacesx.UserData, error) {
	userData, err := us.userRepo.FetchUserDetails(userEmail)
	if err != nil {
		return nil, err
	}
	return &interfacesx.UserData{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		Username:  userData.Username,
		UserRole:  userData.UserRole,
		CreatedAt: userData.CreatedAt,
	}, nil
}
