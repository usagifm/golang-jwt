package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/usagifm/golang-jwt/dto"
	"github.com/usagifm/golang-jwt/entity"
	"github.com/usagifm/golang-jwt/repository"
)

// User service is a contract that defines what users can do
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	MyDiaries(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

// Create a new instance of userservice
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}

}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {

	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}

	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser

}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)

}

func (service *userService) MyDiaries(userID string) entity.User {
	return service.userRepository.MyDiaries(userID)

}
