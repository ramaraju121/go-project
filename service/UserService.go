package service

import (
	"GO_PROJECT/model"
	"GO_PROJECT/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user model.User) (primitive.ObjectID, error) {
	result, err := s.userRepository.CreateUser(user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return id, nil
}

func (s *UserService) GetUser(id primitive.ObjectID) (model.User, error) {
	return s.userRepository.GetUser(id)
}

func (s *UserService) GetAllUser() ([]model.User, error) {
	return s.userRepository.GetAllUser()
}

func (s *UserService) UpdateUser(id primitive.ObjectID, user model.User) error {
	_, err := s.userRepository.UpdateUser(id, user)
	return err
}

func (s *UserService) DeleteUser(id primitive.ObjectID) error {
	_, err := s.userRepository.DeleteUser(id)
	return err
}
