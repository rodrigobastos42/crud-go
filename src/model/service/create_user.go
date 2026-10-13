package service

import (
	"fmt"

	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/configuration/logger"
	"github.com/rodrigobastos42/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *internalerrors.RestErr) {
	logger.Info("CreateUser Model", zap.String("journey", "CreateUser"))

	user.EncryptPassword()

	emailAlreadyRegistered, _ := ud.repository.FindUserByEmail(user.GetEmail())

	if emailAlreadyRegistered != nil {
		return nil, internalerrors.NewBadRequestError(fmt.Sprintf("Email %v already registered", user.GetEmail()))
	}

	createdUser, err := ud.repository.CreateUser(user)

	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
