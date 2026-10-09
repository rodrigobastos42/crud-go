package service

import (
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/model"
)

type UserDomainService interface {
	CreateUser(user model.UserDomainInterface) *internalerrors.RestErr
	FindUser(userId string) (*model.UserDomainInterface, *internalerrors.RestErr)
	UpdateUser(userId string, user model.UserDomainInterface) *internalerrors.RestErr
	DeleteUser(userId string) *internalerrors.RestErr
}

type userDomainService struct {
}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}
