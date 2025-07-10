package service

import "user-service/internal/model"

type UserService interface {
    GetAll() []model.User
}

type userService struct {
    users []model.User
}

func NewUserService() UserService {
    return &userService{
        users: []model.User{
            {ID: 1, Name: "Alice"},
            {ID: 2, Name: "Bob"},
        },
    }
}

func (s *userService) GetAll() []model.User {
    return s.users
}
