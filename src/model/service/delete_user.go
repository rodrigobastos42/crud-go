package service

import (
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(id string) *internalerrors.RestErr {

	logger.Info("DeleteUser Service", zap.String("journey", "DeleteUser"))

	err := ud.repository.DeleteUser(id)

	if err != nil {
		return err
	}

	logger.Info("DeleteUser Service executed successfuly")
	return nil
}
