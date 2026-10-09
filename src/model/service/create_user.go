package service

import (
	"fmt"

	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/configuration/logger"
	"github.com/rodrigobastos42/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(user model.UserDomainInterface) *internalerrors.RestErr {
	logger.Info("CreateUser Model", zap.String("journey", "CreateUser"))

	user.EncryptPassword()

	fmt.Println("USER: ", user)
	return nil
}
