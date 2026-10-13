package service

import (
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/model"
)

func (ud *userDomainService) FindUserByID(id string) (*model.UserDomainInterface, *internalerrors.RestErr) {
	return ud.repository.FindUserByID(id)
}
