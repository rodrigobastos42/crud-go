package service

import (
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/model"
	"github.com/rodrigobastos42/crud-go/src/model/repository"
)

type UserDomainService interface {
	CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *internalerrors.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *internalerrors.RestErr)
	UpdateUser(userId string, user model.UserDomainInterface) *internalerrors.RestErr
	DeleteUser(userId string) *internalerrors.RestErr
	Login(user model.UserDomainInterface) (model.UserDomainInterface, string, *internalerrors.RestErr)
}

type userDomainService struct {
	repository repository.UserRepository
}

func NewUserDomainService(repository *repository.UserRepository) UserDomainService {
	return &userDomainService{
		repository: *repository,
	}
}
