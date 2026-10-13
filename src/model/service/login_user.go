package service

import (
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/configuration/logger"
	"github.com/rodrigobastos42/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) Login(user model.UserDomainInterface) (model.UserDomainInterface, string, *internalerrors.RestErr) {
	logger.Info("Init login service", zap.String("journey", "login"))

	user.EncryptPassword()

	userLoggedIn, err := ud.findUserByEmailAndEncryptPassword(user.GetEmail(), user.GetPassword())

	if err != nil {
		return nil, "", err
	}

	token, err := userLoggedIn.GenerateToken()

	if err != nil {
		return nil, "", err
	}

	logger.Info("login service executed successfully", zap.String("journey", "login"))
	return userLoggedIn, token, nil
}
