package service

import (
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/configuration/logger"
	"github.com/rodrigobastos42/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(id string, user model.UserDomainInterface) *internalerrors.RestErr {

	logger.Info("UpdateUser Service", zap.String("journey", "UpdateUser"))

	err := ud.repository.UpdateUser(id, user)

	if err != nil {
		return err
	}

	logger.Info("UpdateUser Service executed successfuly")
	return nil
}
