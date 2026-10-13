package service

import (
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/model"
)

func (ud *userDomainService) FindUserByID(id string) (model.UserDomainInterface, *internalerrors.RestErr) {
	return ud.repository.FindUserByID(id)
}

func (ud *userDomainService) findUserByEmailAndEncryptPassword(email, password string) (model.UserDomainInterface, *internalerrors.RestErr) {
	return ud.repository.FindUserByEmailAndEncryptPassword(email, password)
}
